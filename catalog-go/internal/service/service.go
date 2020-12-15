package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
)

type Service struct {
	pb.UnimplementedCatalogServiceServer
	listProducts *ListProducts
}

func New(repo Repository) *Service {
	return &Service{listProducts: NewListProducts(repo)}
}

func (s *Service) ListProducts(ctx context.Context, request *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return s.listProducts.Execute(ctx, request)
}
