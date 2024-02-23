package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type ShippingMethods struct {
	gorm.Model

	StoreID uint64                `gorm:"column:store_id; type:bigint(20) unsigned"`
	Method  utils.ShippingMethods `gorm:"column:method; type:tinyint(4)"`
}

func (ShippingMethods) TableName() string {
	return "store_shipping_options"
}
