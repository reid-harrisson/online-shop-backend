package models

import (
	"gorm.io/gorm"
)

type ShippingClasses struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
}

func (ShippingClasses) TableName() string {
	return "store_shipping_classes"
}
