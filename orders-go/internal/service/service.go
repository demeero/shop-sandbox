package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Repository interface {
	Fetch(context.Context, *pb.ListOrdersRequest) ([]*pb.Order, string, error)
	Create(context.Context, *pb.Order) (string, error)
}

type Service struct {
	pb.UnimplementedOrderServiceServer
	listOrders  *ListOrders
	createOrder *CreateOrder
}

func New(repo Repository) *Service {
	return &Service{listOrders: NewListOrders(repo), createOrder: NewCreateOrder(repo)}
}

func (s *Service) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return s.listOrders.Execute(ctx, request)
}

func (s *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return s.createOrder.Execute(ctx, req)
}
