package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/github.com/demeero/shop-sandbox/proto/gen/go/user/v1beta1"
)

var users = map[string]*pb.User{
	"1": {
		Id:    "1",
		Email: "user1@gmail.com",
	},
	"2": {
		Id:    "2",
		Email: "user1@gmail.com",
	},
	"3": {
		Id:    "3",
		Email: "user1@gmail.com",
	},
	"4": {
		Id:    "4",
		Email: "user1@gmail.com",
	},
	"5": {
		Id:    "5",
		Email: "user1@gmail.com",
	},
}

type ListUsers struct {
}

func NewListUsers() *ListUsers {
	return &ListUsers{}
}

func (c *ListUsers) Execute(_ context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	result := make([]*pb.User, 0, len(users))
	for _, u := range users {
		result = append(result, u)
	}
	return &pb.ListUsersResponse{Users: result}, nil
}
