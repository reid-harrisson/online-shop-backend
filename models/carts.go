package models

import "github.com/jinzhu/gorm"

type CartItems struct {
	gorm.Model

	CustomerID  uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	VariationID uint64  `gorm:"column:variation_id; type:bigint(20) unsigned"`
	Quantity    float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

type CartItemsWithDetail struct {
	CartItems

	StoreID     uint64  `gorm:"column:store_id"`
	ProductName string  `gorm:"column:product_name"`
	ImageUrl    string  `gorm:"column:image_url"`
	UnitPrice   float64 `gorm:"column:unit_price"`
	Category    string  `gorm:"column:category"`
	TotalPrice  float64 `gorm:"column:total_price"`
}

type CartItemCount struct {
	Count uint64 `gorm:"column:count"`
}

func (CartItems) TableName() string {
	return "store_cart_items"
}
