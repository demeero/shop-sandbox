package gorm

import (
	"time"
)

type orderStatus struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type order struct {
	ID              string `gorm:"primarykey"`
	UserID          string
	OrderStatusID   int
	CreatedAt       time.Time
	OrderStatus     orderStatus     `gorm:"foreignKey:OrderStatusID"`
	OrderItems      []orderItem     `gorm:"foreignKey:OrderID"`
	ShippingAddress shippingAddress `gorm:"embedded"`
	Total           money           `gorm:"embedded;embeddedPrefix:total_"`
}

type shippingAddress struct {
	ContactName string
	Phone       string
	City        string
	Address1    string
	Address2    string
}

type money struct {
	Units uint
	Nanos uint
}

type orderItem struct {
	ID       string `gorm:"primarykey"`
	OrderID  string
	Product  product `gorm:"embedded;embeddedPrefix:product_"`
	Quantity uint32
	Amount   money `gorm:"embedded;embeddedPrefix:total_"`
}

type product struct {
	ID   string
	Name string
}
