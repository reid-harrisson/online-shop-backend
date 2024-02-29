package models

import (
	"github.com/jinzhu/gorm"
)

type ShippingData struct {
	gorm.Model

	VariationID    uint64  `gorm:"column:variation_id; type:bigint(20) unsigned"`
	Weight         float64 `gorm:"column:weight; type:decimal(20,6)"`
	Width          float64 `gorm:"column:width; type:decimal(20,6)"`
	Height         float64 `gorm:"column:height; type:decimal(20,6)"`
	Length         float64 `gorm:"column:depth; type:decimal(20,6)"`
	Classification string  `gorm:"column:classification; type:varchar(45)"`
}

func (ShippingData) TableName() string {
	return "store_shipping_data"
}
