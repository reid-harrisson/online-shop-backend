package models

import "github.com/jinzhu/gorm"

type ShippingMethods int8

const (
	MethodPickUp ShippingMethods = iota + 1
	MethodFlatRate
	MethodTableRate
	MethodFreeShipping
	MethodRealTime
)

type ShippingOptions struct {
	gorm.Model

	StoreID uint64          `gorm:"column:store_id; type:bigint(20) unsigned"`
	Method  ShippingMethods `gorm:"column:method; type:tinyint(4)"`
}

func (ShippingOptions) TableName() string {
	return "store_shipping_options"
}
