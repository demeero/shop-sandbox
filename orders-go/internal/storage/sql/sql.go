package sql

import (
	"context"
	"database/sql"

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

func (s *Storage) Fetch(ctx context.Context) ([]*orderPb.Order, error) {
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
			   address2
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
		err := rows.Scan(&o.Id, &o.CustomerId, &o.Status, &o.Total.Units, &o.Total.Nanos, &o.ShippingAddress.ContactName,
			&o.ShippingAddress.Phone, &o.ShippingAddress.City, &o.ShippingAddress.Address1, &address2)
		if err != nil {
			return nil, err
		}
		o.ShippingAddress.Address2 = address2.String
		orders = append(orders, o)
	}
	return orders, nil
}
