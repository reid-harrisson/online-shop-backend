package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type ShippingMethods struct {
	gorm.Model

	StoreID       uint64                `gorm:"column:store_id; type:bigint(20) unsigned"`
	Method        utils.ShippingMethods `gorm:"column:method; type:tinyint(4)"`
	FlatRate      float64               `gorm:"column:flat_rate; type:decimal(20,6)"`
	BaseRate      float64               `gorm:"column:base_rate; type:decimal(20,6)"`
	RatePerItem   float64               `gorm:"column:rate_per_item; type:decimal(20,6)"`
	RatePerWeight float64               `gorm:"column:rate_per_weight; type:decimal(20,6)"`
	RatePerTotal  float64               `gorm:"column:rate_per_total; type:decimal(20,6)"`
}

func (ShippingMethods) TableName() string {
	return "store_shipping_methods"
}
