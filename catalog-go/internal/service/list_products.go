package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
)

type ListProducts struct {
	repo Repository
}

func NewListUsers(repo Repository) *ListProducts {
	return &ListProducts{repo: repo}
}

func (c *ListProducts) Execute(ctx context.Context, _ *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := c.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListProductsResponse{Products: products}, nil
}
