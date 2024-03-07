package models

import (
	"gorm.io/gorm"
)

type ShippingTableRates struct {
	gorm.Model

	MethodID    uint64  `gorm:"type:bigint(20)"`
	ClassID     uint64  `gorm:"type:bigint(20)"`
	Condition   int8    `gorm:"type:tinyint(4)"`
	Min         float64 `gorm:"type:decimal(20,6)"`
	Max         float64 `gorm:"type:decimal(20,6)"`
	Break       int8    `gorm:"type:tinyint(4)"`
	Abort       int8    `gorm:"type:tinyint(4)"`
	RowCost     float64 `gorm:"type:decimal(20,6)"`
	ItemCost    float64 `gorm:"type:decimal(20,6)"`
	CostPerKg   float64 `gorm:"type:decimal(20,6)"`
	PercentCost float64 `gorm:"type:decimal(20,6)"`
}

func (ShippingTableRates) TableName() string {
	return "store_shipping_flat_rates"
}
