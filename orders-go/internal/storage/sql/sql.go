package sql

import (
	"context"
	"database/sql"
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

func (s *Storage) Fetch(ctx context.Context, _ *orderPb.ListOrdersRequest) ([]*orderPb.Order, error) {
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

	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*orderPb.Order
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}

		o := &orderPb.Order{}
		o.Total = &moneyPb.Money{}
		o.ShippingAddress = &orderPb.ShippingAddress{}
		var address2 sql.NullString
		var createTime time.Time
		err := rows.Scan(&o.Id, &o.UserId, &o.Status, &o.Total.Units, &o.Total.Nanos, &o.ShippingAddress.ContactName,
			&o.ShippingAddress.Phone, &o.ShippingAddress.City, &o.ShippingAddress.Address1, &address2, &createTime)
		if err != nil {
			return nil, err
		}
		o.ShippingAddress.Address2 = address2.String
		o.CreateTime, _ = ptypes.TimestampProto(createTime)

		orderItems, err := s.fetchOrderItems(ctx, o.Id)
		if err != nil {
			return nil, err
		}

		o.Items = append(o.Items, orderItems...)
		orders = append(orders, o)
	}
	return orders, nil
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
