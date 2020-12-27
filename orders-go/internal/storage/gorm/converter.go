package gorm

import (
	catalogPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
	moneyPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/money/v1"
	orderPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/golang/protobuf/ptypes"
)

func ConvertOrders(orders []order) []*orderPb.Order {
	result := make([]*orderPb.Order, 0, len(orders))
	for _, o := range orders {
		result = append(result, ConvertOrder(o))
	}
	return result
}

func ConvertOrder(order order) *orderPb.Order {
	createTime, _ := ptypes.TimestampProto(order.CreatedAt)
	return &orderPb.Order{
		Id:     order.ID,
		UserId: order.UserID,
		Total: &moneyPb.Money{
			Units: int64(order.Total.Units),
			Nanos: int32(order.Total.Nanos),
		},
		ShippingAddress: &orderPb.ShippingAddress{
			ContactName: order.ShippingAddress.ContactName,
			Phone:       order.ShippingAddress.Phone,
			City:        order.ShippingAddress.City,
			Address1:    order.ShippingAddress.Address1,
			Address2:    order.ShippingAddress.Address2,
		},
		Status:     orderPb.Status(order.OrderStatusID),
		CreateTime: createTime,
		Items:      ConvertOrderItems(order.OrderItems),
	}
}

func ConvertOrderItems(orderItems []orderItem) []*orderPb.OrderItem {
	result := make([]*orderPb.OrderItem, 0, len(orderItems))
	for _, oi := range orderItems {
		result = append(result, ConvertOrderItem(oi))
	}
	return result
}

func ConvertOrderItem(orderItem orderItem) *orderPb.OrderItem {
	return &orderPb.OrderItem{
		Id: orderItem.ID,
		Product: &catalogPb.Product{
			Id:   orderItem.Product.ID,
			Name: orderItem.Product.Name,
		},
		Quantity: orderItem.Quantity,
		Amount: &moneyPb.Money{
			Units: int64(orderItem.Amount.Units),
			Nanos: int32(orderItem.Amount.Nanos),
		},
	}
}
