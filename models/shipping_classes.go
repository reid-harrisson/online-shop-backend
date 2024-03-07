package models

import (
	"gorm.io/gorm"
)

type ShippingClasses struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100)"`
	StoreID     uint64 `gorm:"type:bigint(20)"`
	Description string `gorm:"type:text"`
	Priority    int8   `gorm:"type:tinyint(4)"`
}

func (ShippingClasses) TableName() string {
	return "store_shipping_classes"
}
