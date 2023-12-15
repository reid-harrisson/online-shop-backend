package models

import "github.com/jinzhu/gorm"

type ShippingData struct {
	gorm.Model

	Weight         float64 `gorm:"column:weight; type:decimal(20,6)"`
	Dimension      string  `gorm:"column:dimension; type:varchar(45)"`
	Classification string  `gorm:"column:classification; type:varchar(45)"`
}

func (ShippingData) TableName() string {
	return "shipping_data"
}
