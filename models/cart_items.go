package models

import "github.com/jinzhu/gorm"

type CartItems struct {
	gorm.Model

	ProductID uint64  `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	UserID    uint64  `gorm:"column:user_id; type:bigint(20) unsigned"`
	Quantity  float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

type CartItemDetails struct {
	CartItems
	Product      string  `gorm:"column:product"`
	ProductPrice float64 `gorm:"column:product_price"`
	Store        string  `gorm:"column:store"`
	StoreID      uint64  `gorm:"column:store_id"`
	User         string  `gorm:"column:user"`
	Quantity     float64 `gorm:"column:quantity"`
}

func (CartItems) TableName() string {
	return "store_cart_items"
}
