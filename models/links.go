package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type ProductLinks struct {
	gorm.Model

	ProductID uint64          `gorm:"column:product_id; type:bigint(20) unsigned"`
	LinkID    uint64          `gorm:"column:link_id; type:bigint(20) unsigned"`
	IsUpCross utils.SellTypes `gorm:"column:is_up_cross; type:tinyint(4)"`
}

type ProductsWithLink struct {
	Products
	IsUpCross utils.SellTypes `gorm:"column:is_up_cross"`
}

func (ProductLinks) TableName() string {
	return "store_product_links"
}
