package gorm

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/demeero/shop-sandbox/orders/internal/storage/pagetoken"
)

type Storage struct {
	db *gorm.DB
}

func New(datasource string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(datasource), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Fetch(ctx context.Context, req *pb.ListOrdersRequest) ([]*pb.Order, string, error) {
	q, err := s.buildFetchOrdersQuery(ctx, req)
	if err != nil {
		return nil, "", err
	}
	orders := make([]order, 0, req.GetPageSize())
	if err := q.Find(&orders).Error; err != nil {
		return nil, "", err
	}

	var nextToken string
	if len(orders) == int(req.GetPageSize()) {
		last := orders[len(orders)-1]
		nextToken = pagetoken.PageToken{ID: last.ID, CreatedAt: last.CreatedAt}.Encode()
	}

	return ConvertOrders(orders), nextToken, nil
}

func (s *Storage) buildFetchOrdersQuery(ctx context.Context, req *pb.ListOrdersRequest) (*gorm.DB, error) {
	token := &pagetoken.PageToken{}
	if err := token.Decode(req.GetPageToken()); err != nil {
		return nil, err
	}

	stmnt := s.db.WithContext(ctx).
		Joins("OrderStatus").
		Preload("OrderItems")

	if len(req.GetIds()) > 0 {
		return stmnt.Where(`"order".id IN ?`, req.GetIds()), nil
	}

	// ordering
	orderType := "DESC"
	orderField := "created_at"
	if req.GetOrder() == pb.ListOrdersRequestOrder_LIST_ORDERS_REQUEST_ORDER_ASC {
		orderType = "ASC"
	}
	if req.GetSort() == pb.ListOrdersRequestSort_LIST_ORDERS_REQUEST_SORT_CREATED_AT {
		stmnt = stmnt.Order(fmt.Sprintf("%s %s", orderField, orderType))
	}

	// pagination
	if token.Valid() {
		createdAt := token.CreatedAt.Format(time.RFC3339Nano)
		stmnt = stmnt.Where(`"order".id < ? AND "order".created_at <= ?`, token.ID, createdAt)
	}

	return stmnt.Limit(int(req.GetPageSize())), nil
}

func (s *Storage) Create(_ context.Context, _ *pb.Order) (string, error) {
	panic("implement me")
}

func (s *Storage) UpdateStatus(_ context.Context, _, _ string) (bool, error) {
	panic("implement me")
}
