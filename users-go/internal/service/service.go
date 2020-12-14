package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/github.com/demeero/shop-sandbox/proto/gen/go/user/v1beta1"
)

type Service struct {
	pb.UnimplementedUserAPIServer
	listUsers *ListUsers
}

func New() *Service {
	return &Service{listUsers: NewListUsers()}
}

func (s *Service) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.listUsers.Execute(ctx, request)
}
