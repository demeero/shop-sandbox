package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Service struct {
	pb.UnimplementedOrderServiceServer
	listOrders *ListProducts
}

func New(repo Repository) *Service {
	return &Service{listOrders: NewListOrders(repo)}
}

func (s *Service) ListProducts(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return s.listOrders.Execute(ctx, request)
}
