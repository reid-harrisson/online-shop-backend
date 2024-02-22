package models

import (
	"github.com/jinzhu/gorm"
)

type StoreProductVariations struct {
	gorm.Model

	Sku       string  `gorm:"column:sku; type:varchar(50)"`
	ProductID uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	Price     float64 `gorm:"column:price; type:decimal(20,6)"`
	Quantity  float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

func (StoreProductVariations) TableName() string {
	return "store_product_variations"
}
