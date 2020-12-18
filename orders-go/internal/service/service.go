package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Repository interface {
	Fetch(context.Context) ([]*pb.Order, error)
}

type Service struct {
	pb.UnimplementedOrderServiceServer
	listOrders *ListOrders
}

func New(repo Repository) *Service {
	return &Service{listOrders: NewListOrders(repo)}
}

func (s *Service) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return s.listOrders.Execute(ctx, request)
}
