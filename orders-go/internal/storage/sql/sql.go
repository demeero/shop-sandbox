package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"

	catalogPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
	moneyPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/money/v1"
	orderPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Storage struct {
	db *sql.DB
}

func New(driver, datasource string) (*Storage, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) UpdateStatus(ctx context.Context, statusID, orderID string) (bool, error) {
	var updated bool
	err := Tx(ctx, s.db, func(tx *sql.Tx) error {
		q := `
			UPDATE "order"
			SET order_status_id=$1
			WHERE id = $2
		`
		res, err := tx.ExecContext(ctx, q, statusID, orderID)
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

func (s *Storage) Create(ctx context.Context, order *orderPb.Order) (string, error) {
	var orderID string
	err := Tx(ctx, s.db, func(tx *sql.Tx) error {
		insertOrderQ := `
			INSERT INTO "order" (user_id, order_status_id, total_units, total_nanos, contact_name, phone, city, address1, address2) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id 
		`
		total := order.GetTotal()
		address := order.GetShippingAddress()
		err := tx.QueryRowContext(ctx, insertOrderQ, order.GetUserId(), order.GetStatus(), total.GetUnits(),
			total.GetNanos(), address.GetContactName(), address.GetPhone(), address.GetCity(), address.GetAddress1(), address.GetAddress2()).
			Scan(&orderID)
		if err != nil {
			return err
		}

		insertOrderItemQ := `
			INSERT INTO order_item (quantity, total_units, total_nanos, product_id, product_name, order_id) 
			VALUES ($1, $2, $3, $4, $5, $6)
		`
		for _, oi := range order.GetItems() {
			amount := oi.GetAmount()
			product := oi.GetProduct()
			_, err := tx.ExecContext(ctx, insertOrderItemQ, oi.GetQuantity(), amount.GetUnits(), amount.GetNanos(),
				product.GetId(), product.GetName(), orderID)
			if err != nil {
				return err
			}
		}
		return nil
	}, nil)
	if err != nil {
		return "", err
	}
	return orderID, nil
}

func (s *Storage) Fetch(ctx context.Context, req *orderPb.ListOrdersRequest) ([]*orderPb.Order, string, error) {
	q := buildFetchOrdersQuery(req)
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, "", err
	}
	defer rows.Close()

	var orders []*orderPb.Order
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
	return orders, nextToken, nil
}

func (s *Storage) fetchOrderItems(ctx context.Context, orderID string) ([]*orderPb.OrderItem, error) {
	q := `
		SELECT id,
			   quantity,
			   total_units,
			   total_nanos,
			   product_id,
			   product_name
		FROM order_item
		WHERE order_id = $1  
	`

	rows, err := s.db.QueryContext(ctx, q, orderID)
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

func buildFetchOrdersQuery(req *orderPb.ListOrdersRequest) string {
	q := `
		SELECT id,
			   user_id,
			   order_status_id,
			   total_units,
			   total_nanos,
			   contact_name,
			   phone,
			   city,
			   address1,
			   address2,
			   created_at
		FROM "order"
	`
	// filtering by id
	if len(req.GetIds()) > 0 {
		return q + " WHERE id IN(" + strings.Join(req.GetIds(), ",") + ")"
	}

	// ordering
	orderClause := ""
	if req.GetSort() == orderPb.ListOrdersRequestSort_LIST_ORDERS_REQUEST_SORT_CREATED_AT {
		orderClause = "ORDER BY"
	}
	orderField := "created_at"
	orderType := "DESC"
	if req.GetOrder() == orderPb.ListOrdersRequestOrder_LIST_ORDERS_REQUEST_ORDER_ASC {
		orderType = "ASC"
	}
	orderClause = fmt.Sprintf("%s %s %s", orderClause, orderField, orderType)

	// pagination
	whereClause := "WHERE "
	if req.GetPageToken() == "" {
		whereClause += "id > 0"
	} else {
		whereClause += fmt.Sprintf("id <= '%s'", req.GetPageToken())
	}

	return fmt.Sprintf("%s %s %s %s", q, whereClause, orderClause, "LIMIT "+strconv.Itoa(int(req.GetPageSize()+1)))
}
