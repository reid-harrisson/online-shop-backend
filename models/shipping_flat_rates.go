package models

import (
	"gorm.io/gorm"
)

type ShippingFlatRates struct {
	gorm.Model

	MethodID uint64  `gorm:"type:bigint(20)"`
	ClassID  uint64  `gorm:"type:bigint(20)"`
	BaseCost float64 `gorm:"type:decimal(20,6)"`
	Percent  float64 `gorm:"type:decimal(20,6)"`
	MinFee   float64 `gorm:"type:decimal(20,6)"`
	MaxFee   float64 `gorm:"type:decimal(20,6)"`
}

func (ShippingFlatRates) TableName() string {
	return "shipping_flat_rates"
}
