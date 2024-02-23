package models

import (
	"github.com/jinzhu/gorm"
)

type ProductVariations struct {
	gorm.Model

	ProductID  uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	Sku        string  `gorm:"column:sku; type:varchar(50)"`
	Price      float64 `gorm:"column:price; type:decimal(20,6)"`
	StockLevel float64 `gorm:"column:stock_level; type:decimal(20,6)"`
}

func (ProductVariations) TableName() string {
	return "store_product_variations"
}
