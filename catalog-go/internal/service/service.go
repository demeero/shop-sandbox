package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
)

type Service struct {
	pb.UnimplementedCatalogServiceServer
	listUsers *ListProducts
}

func New(repo Repository) *Service {
	return &Service{listUsers: NewListProducts(repo)}
}

func (s *Service) ListProducts(ctx context.Context, request *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return s.listUsers.Execute(ctx, request)
}
