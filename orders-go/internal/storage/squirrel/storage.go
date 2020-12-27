package squirrel

import (
	"context"
	"database/sql"
	"strings"
	"time"

	catalogPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
	moneyPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/money/v1"
	orderPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/demeero/shop-sandbox/orders/internal/storage/pagetoken"
	"github.com/demeero/shop-sandbox/orders/internal/storage/tx"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes"
)

type Storage struct {
	db   *sql.DB
	psql sq.StatementBuilderType
}

func New(ds string) (*Storage, error) {
	db, err := sql.Open("pgx", ds)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{db: db, psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar)}, nil
}

func (s *Storage) Create(ctx context.Context, o *orderPb.Order) (string, error) {
	total := o.GetTotal()
	addr := o.GetShippingAddress()
	queryOrder := s.psql.Insert(`"order"`).
		Columns("user_id", "order_status_id", "total_units", "total_nanos", "contact_name", "phone", "city",
			"address1", "address2").
		Values(o.GetUserId(), o.GetStatus(), total.GetUnits(), total.GetNanos(), addr.GetContactName(), addr.GetPhone(),
			addr.GetCity(), addr.GetAddress1(), addr.GetAddress2()).
		Suffix("RETURNING id")

	var id string

	queryOrderItem := s.psql.Insert("order_item").
		Columns("quantity", "total_units", "total_nanos", "product_id", "product_name", "order_id")
	for _, oi := range o.GetItems() {
		a := oi.GetAmount()
		p := oi.GetProduct()
		queryOrderItem = queryOrderItem.Values(oi.GetQuantity(), a.GetUnits(), a.GetNanos(), p.GetId(), p.GetName(), &id)
	}

	err := tx.Tx(ctx, s.db, func(tx *sql.Tx) error {
		if err := queryOrder.RunWith(tx).QueryRowContext(ctx).Scan(&id); err != nil {
			return err
		}
		_, err := queryOrderItem.RunWith(tx).ExecContext(ctx)
		return err
	}, nil)
	return id, err
}

func (s *Storage) UpdateStatus(ctx context.Context, statusID, orderID string) (bool, error) {
	var updated bool
	err := tx.Tx(ctx, s.db, func(tx *sql.Tx) error {
		res, err := s.psql.Update(`"order"`).
			Where(sq.Eq{"id": orderID}).
			Set("order_status_id", statusID).
			RunWith(tx).
			ExecContext(ctx)
		if err != nil {
			return err
		}
		n, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if n > 0 {
			updated = true
		}
		return nil
	}, nil)
	return updated, err
}

func (s *Storage) Fetch(ctx context.Context, req *orderPb.ListOrdersRequest) ([]*orderPb.Order, string, error) {
	rows, err := s.fetchOrders(ctx, req)
	if err != nil {
		return nil, "", err
	}
	defer rows.Close()

	orders := make([]*orderPb.Order, 0, req.GetPageSize())
	var nextToken string
	for i := 0; rows.Next(); i++ {
		if err := rows.Err(); err != nil {
			return nil, "", err
		}

		o := &orderPb.Order{}
		o.Total = &moneyPb.Money{}
		o.ShippingAddress = &orderPb.ShippingAddress{}
		var address2 sql.NullString
		var createTime time.Time
		err := rows.Scan(&o.Id, &o.UserId, &o.Status, &o.Total.Units, &o.Total.Nanos, &o.ShippingAddress.ContactName,
			&o.ShippingAddress.Phone, &o.ShippingAddress.City, &o.ShippingAddress.Address1, &address2, &createTime)
		if err != nil {
			return nil, "", err
		}
		o.ShippingAddress.Address2 = address2.String
		o.CreateTime, _ = ptypes.TimestampProto(createTime)

		orderItems, err := s.fetchOrderItems(ctx, o.Id)
		if err != nil {
			return nil, "", err
		}

		if i == int(req.GetPageSize()) {
			nextToken = o.GetId()
			break
		}

		o.Items = append(o.Items, orderItems...)
		orders = append(orders, o)
	}
	ordersLen := len(orders)
	if ordersLen == int(req.GetPageSize()) {
		last := orders[ordersLen-1]
		createTime, _ := ptypes.Timestamp(last.GetCreateTime())
		nextToken = pagetoken.PageToken{ID: last.GetId(), CreatedAt: createTime}.Encode()
	}
	return orders, nextToken, nil
}

func (s *Storage) fetchOrderItems(ctx context.Context, orderID string) ([]*orderPb.OrderItem, error) {
	rows, err := s.psql.Select("id", "quantity", "total_units", "total_nanos", "product_id", "product_name").
		From("order_item").
		Where(sq.Eq{"order_id": orderID}).
		RunWith(s.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderItems []*orderPb.OrderItem
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		oi := &orderPb.OrderItem{}
		oi.Amount = &moneyPb.Money{}
		oi.Product = &catalogPb.Product{}
		var productName string
		err := rows.Scan(&oi.Id, &oi.Quantity, &oi.Amount.Units, &oi.Amount.Nanos, &oi.Product.Id, &productName)
		if err != nil {
			return nil, err
		}
		oi.Product.Name = productName

		orderItems = append(orderItems, oi)
	}
	return orderItems, nil
}

func (s *Storage) fetchOrders(ctx context.Context, req *orderPb.ListOrdersRequest) (*sql.Rows, error) {
	token := &pagetoken.PageToken{}
	if err := token.Decode(req.GetPageToken()); err != nil {
		return nil, err
	}

	q := s.psql.Select("id", "user_id", "order_status_id", "total_units", "total_nanos", "contact_name",
		"phone", "city", "address1", "address2", "created_at").
		From(`"order"`)

	// filtering by id
	if len(req.GetIds()) > 0 {
		return q.Where(sq.Eq{"id": []string{strings.Join(req.GetIds(), ",")}}).RunWith(s.db).QueryContext(ctx)
	}

	// ordering
	orderField := "created_at"
	orderType := "DESC"
	if req.GetOrder() == orderPb.ListOrdersRequestOrder_LIST_ORDERS_REQUEST_ORDER_ASC {
		orderType = "ASC"
	}
	if req.GetSort() == orderPb.ListOrdersRequestSort_LIST_ORDERS_REQUEST_SORT_CREATED_AT {
		q = q.OrderByClause("? "+orderType, orderField)
	}

	// pagination
	if token.Valid() {
		q = q.Where(sq.And{sq.Lt{"id": token.ID}, sq.LtOrEq{"created_at": token.CreatedAt.Format(time.RFC3339Nano)}})
	}
	q = q.Limit(uint64(req.GetPageSize()))

	return q.RunWith(s.db).QueryContext(ctx)
}
