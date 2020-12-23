package service

import (
	"context"

	pb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"
)

type ListOrders struct {
	repo Repository
}

func NewListOrders(repo Repository) *ListOrders {
	return &ListOrders{repo: repo}
}

func (c *ListOrders) Execute(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	setDefaults(req)
	orders, nextToken, err := c.repo.Fetch(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ListOrdersResponse{Orders: orders, NextPageToken: nextToken}, nil
}

func setDefaults(req *pb.ListOrdersRequest) {
	if req.GetOrder() == pb.ListOrdersRequestOrder_LIST_ORDERS_REQUEST_ORDER_UNSPECIFIED {
		req.Order = pb.ListOrdersRequestOrder_LIST_ORDERS_REQUEST_ORDER_DESC
	}
	if req.GetSort() == pb.ListOrdersRequestSort_LIST_ORDERS_REQUEST_SORT_UNSPECIFIED {
		req.Sort = pb.ListOrdersRequestSort_LIST_ORDERS_REQUEST_SORT_CREATED_AT
	}
	if req.GetPageSize() == 0 {
		req.PageSize = 200
	}
}
