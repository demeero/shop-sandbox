package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type Repository interface {
	Fetch(context.Context, *pb.ListOrdersRequest) ([]*pb.Order, string, error)
	Create(context.Context, *pb.Order) (string, error)
	UpdateStatus(ctx context.Context, statusID, orderID string) (bool, error)
}

type Service struct {
	pb.UnimplementedOrderServiceServer
	listOrders   *ListOrders
	createOrder  *CreateOrder
	updateStatus *UpdateStatus
}

func New(repo Repository) *Service {
	return &Service{
		listOrders:   NewListOrders(repo),
		createOrder:  NewCreateOrder(repo),
		updateStatus: NewUpdateStatus(repo),
	}
}

func (s *Service) ListOrders(ctx context.Context, request *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	return s.listOrders.Execute(ctx, request)
}

func (s *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return s.createOrder.Execute(ctx, req)
}

func (s *Service) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	return s.updateStatus.Execute(ctx, req)
}
