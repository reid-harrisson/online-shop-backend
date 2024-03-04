package models

import "gorm.io/gorm"

type ProductAttributeValues struct {
	gorm.Model

	AttributeID    uint64 `gorm:"column:attribute_id; type:bigint(20) unsigned"`
	AttributeValue string `gorm:"column:attribute_value; type:varchar(50)"`
}

func (ProductAttributeValues) TableName() string {
	return "store_product_attribute_values"
}

type ProductAttributeValuesWithDetail struct {
	ProductAttributeValues
	AttributeName string `gorm:"column:attribute_name"`
	Unit          string `gomr:"column:attribute_unit"`
}
