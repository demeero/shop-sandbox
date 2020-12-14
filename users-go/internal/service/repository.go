package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"
)

type Repository interface {
	Fetch(context.Context) ([]*pb.User, error)
}
