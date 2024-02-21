package models

import "github.com/jinzhu/gorm"

type IsUpCross int8

const (
	UpSell IsUpCross = iota
	CrossSell
)

type ProductLinked struct {
	gorm.Model

	ProductID uint64    `gorm:"column:product_id; type:bigint(20) unsigned"`
	LinkedID  uint64    `gorm:"column:linked_id; type:bigint(20) unsigned"`
	IsUpCross IsUpCross `gorm:"column:is_up_cross; type:tinyint(4)"`
}

type ProductsWithLink struct {
	Products
	IsUpCross IsUpCross `gorm:"column:is_up_cross"`
}

func (ProductLinked) TableName() string {
	return "store_linked_products"
}
