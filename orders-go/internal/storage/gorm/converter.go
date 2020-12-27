package gorm

import (
	catalogPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
	moneyPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/money/v1"
	orderPb "github.com/demeero/shop-sandbox/proto/gen/go/shop/order/v1beta1"

	"github.com/golang/protobuf/ptypes"
)

func convertToExternalOrders(orders []order) []*orderPb.Order {
	result := make([]*orderPb.Order, 0, len(orders))
	for _, o := range orders {
		result = append(result, convertToExternalOrder(o))
	}
	return result
}

func convertToExternalOrder(order order) *orderPb.Order {
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
		Status:     orderPb.Status(order.OrderStatus.ID),
		CreateTime: createTime,
		Items:      convertToExternalOrderItems(order.OrderItems),
	}
}

func convertToExternalOrderItems(orderItems []orderItem) []*orderPb.OrderItem {
	result := make([]*orderPb.OrderItem, 0, len(orderItems))
	for _, oi := range orderItems {
		result = append(result, convertToExternalOrderItem(oi))
	}
	return result
}

func convertToExternalOrderItem(orderItem orderItem) *orderPb.OrderItem {
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

func convertToInternalOrder(o *orderPb.Order) order {
	address := o.GetShippingAddress()
	total := o.GetTotal()
	return order{
		UserID:        o.GetUserId(),
		OrderStatusID: int(o.GetStatus()),
		OrderItems:    convertToInternalOrderItems(o.GetItems()),
		ShippingAddress: shippingAddress{
			ContactName: address.GetContactName(),
			Phone:       address.GetPhone(),
			City:        address.GetCity(),
			Address1:    address.GetAddress1(),
			Address2:    address.GetAddress2(),
		},
		Total: money{
			Units: uint(total.GetUnits()),
			Nanos: uint(total.GetNanos()),
		},
	}
}

func convertToInternalOrderItems(orderItems []*orderPb.OrderItem) []orderItem {
	result := make([]orderItem, 0, len(orderItems))
	for _, oi := range orderItems {
		result = append(result, convertToInternalOrderItem(oi))
	}
	return result
}

func convertToInternalOrderItem(oi *orderPb.OrderItem) orderItem {
	return orderItem{
		Product: product{
			ID:   oi.GetProduct().GetId(),
			Name: oi.GetProduct().GetName(),
		},
		Quantity: oi.GetQuantity(),
		Amount: money{
			Units: uint(oi.GetAmount().GetUnits()),
			Nanos: uint(oi.GetAmount().GetNanos()),
		},
	}
}
