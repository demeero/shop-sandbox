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

func (c *CreateOrder) Execute(_ context.Context, _ *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return nil, errors.New("unimplemented")
}
