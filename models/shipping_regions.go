package models

import (
	"gorm.io/gorm"
)

type ShippingRegions struct {
	gorm.Model

	ZoneID    uint64 `gorm:"type:bigint(20)"`
	CountryID uint64 `gorm:"type:bigint(20)"`
	RegionID  uint64 `gorm:"type:bigint(20)"`
	CityID    uint64 `gorm:"type:bigint(20)"`
}

func (ShippingRegions) TableName() string {
	return "store_shipping_regions"
}
