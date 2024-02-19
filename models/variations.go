package models

import "github.com/jinzhu/gorm"

type ProductVariations struct {
	gorm.Model

	AttributeID uint64 `gorm:"column:attribute_id; type:bigint(20) unsigned"`
	Variant     string `gorm:"column:variant; type:varchar(50)"`
}

func (ProductVariations) TableName() string {
	return "store_product_variations"
}

type ProductVariationsWithName struct {
	ProductVariations
	ProductID     uint64 `gorm:"column:product_id"`
	AttributeName string `gorm:"column:attribute_name"`
	AttributeUnit string `gomr:"column:attribute_unit"`
}
