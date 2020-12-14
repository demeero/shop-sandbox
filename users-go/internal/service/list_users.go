package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"
)

type ListUsers struct {
	repo Repository
}

func NewListUsers(repo Repository) *ListUsers {
	return &ListUsers{repo: repo}
}

func (c *ListUsers) Execute(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := c.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListUsersResponse{Users: users}, nil
}
