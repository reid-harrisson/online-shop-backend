package models

import "github.com/jinzhu/gorm"

type ProductAttributeValues struct {
	gorm.Model

	AttributeID uint64 `gorm:"column:attribute_id; type:bigint(20) unsigned"`
	Value       string `gorm:"column:value; type:varchar(50)"`
}

func (ProductAttributeValues) TableName() string {
	return "store_product_attribute_values"
}

type ProductAttributeValuesWithDetail struct {
	ProductAttributeValues
	ProductID     uint64 `gorm:"column:product_id"`
	AttributeName string `gorm:"column:attribute_name"`
	AttributeUnit string `gomr:"column:attribute_unit"`
}
