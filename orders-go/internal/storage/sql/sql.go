package sql

import (
	"context"
	"database/sql"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
	_ "github.com/jackc/pgx/v4/stdlib"
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

func (s *Storage) Fetch(context.Context) ([]*pb.Order, error) {
	return nil, nil
}
