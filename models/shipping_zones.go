package models

import (
	"gorm.io/gorm"
)

type ShippingZones struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100)"`
	StoreID     uint64 `gorm:"type:bigint(20)"`
	Description string `gorm:"type:text"`
}

type ShippingZonesWithPlace struct {
	ShippingZones
	PlaceIDs   string
	PlaceNames string
}

func (ShippingZones) TableName() string {
	return "store_shipping_zones"
}
