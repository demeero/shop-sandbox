package service

import (
	"context"
	"errors"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type CreateOrder struct {
	repo Repository
}

func NewCreateOrder(repo Repository) *CreateOrder {
	return &CreateOrder{repo: repo}
}

// TODO calculate order total
func (c *CreateOrder) Execute(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	if req.GetOrder() == nil {
		return nil, errors.New("field `order` is required")
	}
	req.GetOrder().Status = pb.Status_STATUS_PENDING
	id, err := c.repo.Create(ctx, req.GetOrder())
	if err != nil {
		return nil, err
	}
	result, _, err := c.repo.Fetch(ctx, &pb.ListOrdersRequest{Ids: []string{id}, PageSize: 1})
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("created order not found")
	}
	return &pb.CreateOrderResponse{Order: result[0]}, nil
}
