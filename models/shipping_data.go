package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type ShippingData struct {
	gorm.Model

	ProductID      uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	Weight         float64 `gorm:"column:weight; type:decimal(20,6)"`
	Width          float64 `gorm:"column:width; type:decimal(20,6)"`
	Height         float64 `gorm:"column:height; type:decimal(20,6)"`
	Depth          float64 `gorm:"column:depth; type:decimal(20,6)"`
	Classification string  `gorm:"column:classification; type:varchar(45)"`
}

type ShippingMethods struct {
	gorm.Model

	StoreID uint64                `gorm:"column:store_id; type:bigint(20) unsigned"`
	Method  utils.ShippingMethods `gorm:"column:method; type:tinyint(4)"`
}

func (ShippingData) TableName() string {
	return "store_shipping_data"
}

func (ShippingMethods) TableName() string {
	return "store_shipping_methods"
}
