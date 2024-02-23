package models

import "github.com/jinzhu/gorm"

type ProductVariations struct {
	gorm.Model

	ProductID  uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	SKU        string  `gorm:"column:sku; type:varchar(50)"`
	Price      float64 `gorm:"column:price; type:decimal(20,6)"`
	StockLevel float64 `gorm:"column:stock_level; type:decimal(20,6)"`
}

type ProductVariationDetails struct {
	gorm.Model

	VariationID      uint64 `gorm:"column:variation_id; type:bigint(20) unsigned"`
	AttributeValueID uint64 `gorm:"column:attribute_value_id; type:bigint(20) unsigned"`
}

func (ProductVariations) TableName() string {
	return "store_product_variations"
}

func (ProductVariationDetails) TableName() string {
	return "store_product_variation_details"
}

type ProductVariationsWithName struct {
	ProductVariations

	Title          uint64 `gorm:"column:product_id"`
	AttributeName  string `gorm:"column:attribute_name"`
	AttributeValue string `gomr:"column:attribute_value"`
}
