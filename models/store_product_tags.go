package models

import "github.com/jinzhu/gorm"

type Tags struct {
	gorm.Model

	ProductID uint64 `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	Tag       string `gorm:"column:tag; type:varchar(20)"`
}

func (Tags) TableName() string {
	return "store_product_tags"
}
