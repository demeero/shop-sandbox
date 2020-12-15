package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Repository interface {
	Fetch(context.Context) ([]*pb.Order, error)
}
