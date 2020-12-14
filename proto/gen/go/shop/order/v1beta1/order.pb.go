// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: shop/order/v1beta1/order.proto

package v1beta1

import (
	v1beta1 "github.com/demeero/shop-sandbox/proto/gen/go/shop/catalog/v1beta1"
	v1 "github.com/demeero/shop-sandbox/proto/gen/go/shop/money/v1"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_PENDING     Status = 1
	Status_STATUS_PROCESSING  Status = 2
	Status_STATUS_CANCELED    Status = 3
	Status_STATUS_COMPLETED   Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_PENDING",
		2: "STATUS_PROCESSING",
		3: "STATUS_CANCELED",
		4: "STATUS_COMPLETED",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_PENDING":     1,
		"STATUS_PROCESSING":  2,
		"STATUS_CANCELED":    3,
		"STATUS_COMPLETED":   4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_shop_order_v1beta1_order_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_shop_order_v1beta1_order_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_shop_order_v1beta1_order_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId      string           `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Items           []*OrderItem     `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Total           *v1.Money        `protobuf:"bytes,4,opt,name=total,proto3" json:"total,omitempty"`
	ShippingAddress *ShippingAddress `protobuf:"bytes,5,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	Status          Status           `protobuf:"varint,6,opt,name=status,proto3,enum=shop.order.v1beta1.Status" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_order_v1beta1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_shop_order_v1beta1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_shop_order_v1beta1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetTotal() *v1.Money {
	if x != nil {
		return x.Total
	}
	return nil
}

func (x *Order) GetShippingAddress() *ShippingAddress {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

func (x *Order) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product  *v1beta1.Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Quantity uint32           `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount   *v1.Money        `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_order_v1beta1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_shop_order_v1beta1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_shop_order_v1beta1_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderItem) GetProduct() *v1beta1.Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *OrderItem) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetAmount() *v1.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

type ShippingAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContactName string `protobuf:"bytes,1,opt,name=contact_name,json=contactName,proto3" json:"contact_name,omitempty"`
	Phone       string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Address1    string `protobuf:"bytes,4,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2    string `protobuf:"bytes,5,opt,name=address2,proto3" json:"address2,omitempty"`
}

func (x *ShippingAddress) Reset() {
	*x = ShippingAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_order_v1beta1_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingAddress) ProtoMessage() {}

func (x *ShippingAddress) ProtoReflect() protoreflect.Message {
	mi := &file_shop_order_v1beta1_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingAddress.ProtoReflect.Descriptor instead.
func (*ShippingAddress) Descriptor() ([]byte, []int) {
	return file_shop_order_v1beta1_order_proto_rawDescGZIP(), []int{2}
}

func (x *ShippingAddress) GetContactName() string {
	if x != nil {
		return x.ContactName
	}
	return ""
}

func (x *ShippingAddress) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ShippingAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ShippingAddress) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *ShippingAddress) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

var File_shop_order_v1beta1_order_proto protoreflect.FileDescriptor

var file_shop_order_v1beta1_order_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x22, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x4e, 0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x31, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x2a, 0x76, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6d, 0x65, 0x65, 0x72, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70,
	0x2d, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_order_v1beta1_order_proto_rawDescOnce sync.Once
	file_shop_order_v1beta1_order_proto_rawDescData = file_shop_order_v1beta1_order_proto_rawDesc
)

func file_shop_order_v1beta1_order_proto_rawDescGZIP() []byte {
	file_shop_order_v1beta1_order_proto_rawDescOnce.Do(func() {
		file_shop_order_v1beta1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_order_v1beta1_order_proto_rawDescData)
	})
	return file_shop_order_v1beta1_order_proto_rawDescData
}

var file_shop_order_v1beta1_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shop_order_v1beta1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shop_order_v1beta1_order_proto_goTypes = []interface{}{
	(Status)(0),             // 0: shop.order.v1beta1.Status
	(*Order)(nil),           // 1: shop.order.v1beta1.Order
	(*OrderItem)(nil),       // 2: shop.order.v1beta1.OrderItem
	(*ShippingAddress)(nil), // 3: shop.order.v1beta1.ShippingAddress
	(*v1.Money)(nil),        // 4: shop.money.v1.Money
	(*v1beta1.Product)(nil), // 5: shop.catalog.v1beta1.Product
}
var file_shop_order_v1beta1_order_proto_depIdxs = []int32{
	2, // 0: shop.order.v1beta1.Order.items:type_name -> shop.order.v1beta1.OrderItem
	4, // 1: shop.order.v1beta1.Order.total:type_name -> shop.money.v1.Money
	3, // 2: shop.order.v1beta1.Order.shipping_address:type_name -> shop.order.v1beta1.ShippingAddress
	0, // 3: shop.order.v1beta1.Order.status:type_name -> shop.order.v1beta1.Status
	5, // 4: shop.order.v1beta1.OrderItem.product:type_name -> shop.catalog.v1beta1.Product
	4, // 5: shop.order.v1beta1.OrderItem.amount:type_name -> shop.money.v1.Money
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shop_order_v1beta1_order_proto_init() }
func file_shop_order_v1beta1_order_proto_init() {
	if File_shop_order_v1beta1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_order_v1beta1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shop_order_v1beta1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shop_order_v1beta1_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shop_order_v1beta1_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_order_v1beta1_order_proto_goTypes,
		DependencyIndexes: file_shop_order_v1beta1_order_proto_depIdxs,
		EnumInfos:         file_shop_order_v1beta1_order_proto_enumTypes,
		MessageInfos:      file_shop_order_v1beta1_order_proto_msgTypes,
	}.Build()
	File_shop_order_v1beta1_order_proto = out.File
	file_shop_order_v1beta1_order_proto_rawDesc = nil
	file_shop_order_v1beta1_order_proto_goTypes = nil
	file_shop_order_v1beta1_order_proto_depIdxs = nil
}
