package models

import (
	"gorm.io/gorm"
)

type VariationDetails struct {
	gorm.Model

	VariationID      uint64 `gorm:"column:variation_id; type:bigint(20) unsigned"`
	AttributeValueID uint64 `gorm:"column:attribute_value_id; type:bigint(20) unsigned"`
}

func (VariationDetails) TableName() string {
	return "store_product_variation_details"
}
