package models

import "github.com/jinzhu/gorm"

type Attributes struct {
	gorm.Model

	ProductID uint64 `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	Attribute string `gorm:"column:attribute; type:varchar(20)"`
	Value     string `gorm:"column:value; type:varchar(20)"`
}

func (Attributes) TableName() string {
	return "store_product_attributes"
}
