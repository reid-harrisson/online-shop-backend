package models

import (
	"github.com/jinzhu/gorm"
)

type OrderStatuses int8

const (
	StatusOrderPending OrderStatuses = iota + 1
	StatusOrderPaymentProcessing
	StatusOrderPaid
	StatusOrderProcessing
	StatusOrderShippingProcessing
	StatusOrderShipping
	StatusOrderShipped
	StatusOrderCompleted
)

func OrderStatusToString(orderStatus OrderStatuses) string {
	status := ""
	switch orderStatus {
	case StatusOrderPending:
		status = "Pending"
	case StatusOrderPaymentProcessing:
		status = "Payment Processing"
	case StatusOrderPaid:
		status = "Paid"
	case StatusOrderProcessing:
		status = "Processing"
	case StatusOrderShippingProcessing:
		status = "Shipping Processing"
	case StatusOrderShipping:
		status = "Shipping"
	case StatusOrderShipped:
		status = "Shipped"
	case StatusOrderCompleted:
		status = "Completed"
	}
	return status
}

type Orders struct {
	gorm.Model

	CustomerID      uint64        `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Status          OrderStatuses `gorm:"column:status; type:tinyint(4)"`
	BillingAddress  string        `gorm:"column:billing_address; type:varchar(100)"`
	ShippingAddress string        `gorm:"column:shipping_address; type:varchar(100)"`
}

type OrderItems struct {
	gorm.Model

	OrderID        uint64        `gorm:"column:order_id; type:bigint(20) unsigned"`
	StoreID        uint64        `gorm:"column:store_id; type:bigint(20) unsigned"`
	ProductID      uint64        `gorm:"column:product_id; type:bigint(20) unsigned"`
	UnitPriceSale  float64       `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	Quantity       float64       `gorm:"column:quantity; type:decimal(20,6)"`
	SubTotalPrice  float64       `gorm:"column:sub_total_price; type:decimal(20,6)"`
	TaxRate        float64       `gorm:"column:tax_rate; type:decimal(20,6)"`
	TaxAmount      float64       `gorm:"column:tax_amount; type:decimal(20,6)"`
	ShippingMethod string        `gorm:"column:shipping_method; type:varchar(20)"`
	ShippingPrice  float64       `gorm:"column:shipping_price; type:deciaml(20,6)"`
	TotalPrice     float64       `gorm:"column:total_price; type:decimal(20,6)"`
	Status         OrderStatuses `gorm:"column:status; type:tinyint(4)"`
}

type CustomerOrders struct {
	OrderID         uint64        `gorm:"column:order_id"`
	OrderStatus     OrderStatuses `gorm:"column:order_status"`
	TotalPrice      float64       `gorm:"column:total_price"`
	BillingAddress  string        `gorm:"column:billing_address"`
	ShippingAddress string        `gorm:"column:shipping_address"`
}

type CustomerOrdersWithDetail struct {
	OrderStatus    OrderStatuses `gorm:"column:order_status"`
	StoreID        uint64        `gorm:"column:store_id"`
	ProductStatus  OrderStatuses `gorm:"column:product_status"`
	ProductID      uint64        `gorm:"column:product_id"`
	UnitPriceSale  float64       `gorm:"column:unit_price_sale"`
	Quantity       float64       `gorm:"column:quantity"`
	SubTotalPrice  float64       `gorm:"column:sub_total_price"`
	TaxRate        float64       `gorm:"column:tax_rate"`
	TaxAmount      float64       `gorm:"column:tax_amount"`
	ShippingMethod string        `gorm:"column:shipping_method"`
	ShippingPrice  float64       `gorm:"column:shipping_price"`
	TotalPrice     float64       `gorm:"column:total_price"`
}

type StoreOrders struct {
	OrderID         uint64        `gorm:"column:order_id"`
	CustomerID      uint64        `gorm:"column:customer_id"`
	ProductID       uint64        `gorm:"column:product_id"`
	UnitPriceSale   float64       `gorm:"column:unit_price_sale"`
	Quantity        float64       `gorm:"column:quantity"`
	SubTotalPrice   float64       `gorm:"column:sub_total_price"`
	BillingAddress  string        `gorm:"column:billing_address"`
	ShippingAddress string        `gorm:"column:shipping_address"`
	TaxRate         float64       `gorm:"column:tax_rate"`
	TaxAmount       float64       `gorm:"column:tax_amount"`
	ShippingMethod  string        `gorm:"column:shipping_method"`
	ShippingPrice   float64       `gorm:"column:shipping_price"`
	TotalPrice      float64       `gorm:"column:total_price"`
	ProductStatus   OrderStatuses `gorm:"column:product_status"`
}

func (Orders) TableName() string {
	return "store_orders"
}

func (OrderItems) TableName() string {
	return "store_order_items"
}
