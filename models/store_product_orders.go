package models

import "github.com/jinzhu/gorm"

type ProductOrders struct {
	gorm.Model

	StoreID       uint64  `gorm:"column:store_id; type:bigint(20) unsigned"`
	CustomerID    uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	ProductID     uint64  `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	UnitPriceSale float64 `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	Quantity      float64 `gorm:"column:quantity; type:decimal(20,6)"`
	SubTotal      float64 `gorm:"column:sub_total; type:decimal(20,6)"`
	TaxRate       float64 `gorm:"column:tax_rate; type:decimal(20,6)"`
	TaxAmount     float64 `gorm:"column:tax_amount; type:decimal(20,6)"`
	Total         float64 `gorm:"column:total; type:decimal(20,6)"`
	Status        string  `gorm:"column:status; type:varchar(20)"`
}

func (ProductOrders) TableName() string {
	return "store_product_orders"
}
