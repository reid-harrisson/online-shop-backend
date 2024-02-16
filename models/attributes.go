package models

import "github.com/jinzhu/gorm"

type ProductAttributes struct {
	gorm.Model

	ProductID uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	Name      string `gorm:"column:name; type:varchar(50)"`
	Unit      string `gorm:"column:unit; type:varchar(50)"`
}

func (ProductAttributes) TableName() string {
	return "store_product_attributes"
}
