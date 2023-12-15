package models

import "github.com/jinzhu/gorm"

type CartItems struct {
	gorm.Model

	ProductID uint64  `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	UserID    uint64  `gorm:"column:user_id; type:bigint(20) unsigned"`
	Quantity  float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

func (CartItems) TableName() string {
	return "store_cart_items"
}
