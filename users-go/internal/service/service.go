package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/user/v1beta1"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	listUsers *ListUsers
}

func New(repo Repository) *Service {
	return &Service{listUsers: NewListUsers(repo)}
}

func (s *Service) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.listUsers.Execute(ctx, request)
}
