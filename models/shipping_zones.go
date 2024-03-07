package models

import (
	"gorm.io/gorm"
)

type ShippingZones struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
}

func (ShippingZones) TableName() string {
	return "store_shipping_zones"
}
