package models

import "gorm.io/gorm"

type ProductAttributes struct {
	gorm.Model

	ProductID     uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	AttributeName string `gorm:"column:attribute_name; type:varchar(50)"`
}

func (ProductAttributes) TableName() string {
	return "store_product_attributes"
}
