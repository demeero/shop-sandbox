package service

import (
	"context"
	"errors"
	"strconv"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type UpdateStatus struct {
	repo Repository
}

func NewUpdateStatus(repo Repository) *UpdateStatus {
	return &UpdateStatus{repo: repo}
}

func (c *UpdateStatus) Execute(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	if req.GetStatus() == pb.Status_STATUS_UNSPECIFIED {
		return nil, errors.New("order status should be set")
	}
	ok, err := c.repo.UpdateStatus(ctx, strconv.Itoa(int(req.GetStatus())), req.GetOrderId())
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("not found")
	}
	return &pb.UpdateStatusResponse{}, nil
}
