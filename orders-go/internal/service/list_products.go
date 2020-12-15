package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type ListProducts struct {
	repo Repository
}

func NewListOrders(repo Repository) *ListProducts {
	return &ListProducts{repo: repo}
}

func (c *ListProducts) Execute(ctx context.Context, _ *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := c.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListOrdersResponse{Orders: orders}, nil
}
