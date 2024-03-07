package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type ShippingMethods struct {
	gorm.Model

	ZoneID              uint64                `gorm:"type:bigint(20)"`
	StoreID             uint64                `gorm:"type:bigint(20)"`
	Method              utils.ShippingMethods `gorm:"type:tinyint(4)"`    //Free Shipping | Local Pickup | Flat Rate | Table Rate
	Requirement         int8                  `gorm:"type:tinyint(4)"`    //Free Shipping
	MinimumOrderAmount  float64               `gorm:"type:decimal(20,6)"` //Free Shipping
	TaxStatus           int8                  `gorm:"type:tinyint(4)"`    //Flat Rate, Local Pickup
	Cost                float64               `gorm:"type:decimal(20,6)"` //Flat Rate, Local Pickup
	TaxIncluded         int8                  `gorm:"type:tinyint(4)"`    //Table Rate
	HandlingFee         float64               `gorm:"type:decimal(20,6)"` //Table Rate
	MaximumShippingCost float64               `gorm:"type:decimal(20,6)"` //Table Rate
	CalculationType     int8                  `gorm:"type:tinyint(4)"`    //Table Rate
	HandlingFeePerClass float64               `gorm:"type:decimal(20,6)"` //Table Rate
	MinimumCostPerClass float64               `gorm:"type:decimal(20,6)"` //Table Rate
	MaximumCostPerClass float64               `gorm:"type:decimal(20,6)"` //Table Rate
	DiscountInMinMax    int8                  `gorm:"type:tinyint(4)"`    //Table Rate
	TaxInMinMax         int8                  `gorm:"type:tinyint(4)"`    //Table Rate
}

func (ShippingMethods) TableName() string {
	return "store_shipping_methods"
}
