package models

import "github.com/jinzhu/gorm"

type BaseAttributes struct {
	gorm.Model

	Name string `gorm:"column:name; type:varchar(50)"`
	Unit string `gorm:"column:unit; type:varchar(50)"`
}

type ProductAttributes struct {
	gorm.Model

	ProductID   uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	AttributeID uint64 `gorm:"column:attribute_id; type:bigint(20) unsigned"`
}

func (BaseAttributes) TableName() string {
	return "store_attributes"
}

func (ProductAttributes) TableName() string {
	return "store_product_attributes"
}

type ProductAttributesWithName struct {
	ProductAttributes
	AttributeName string `gorm:"column:attribute_name"`
	AttributeUnit string `gomr:"column:attribute_unit"`
}
