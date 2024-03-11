package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupons struct {
	gorm.Model

	StoreID           uint64    `gorm:"column:store_id; type:bigint(20)"`
	CouponCode        string    `gorm:"column:coupon_code; type:varchar(100)"`
	DiscountType      int8      `gorm:"column:discount_type; type:tinyint(20)"`
	CouponAmount      float64   `gorm:"column:coupon_amount; type:decimal(20,6)"`
	AllowFreeShipping int8      `gorm:"column:allow_free_shipping; type:tinyint(20)"`
	ExpiryDate        time.Time `gorm:"column:expiry_date; type:datetime"`
	MinimumSpend      float64   `gorm:"column:minimum_spend; type:decimal(20,6)"`
	MaximumSpend      float64   `gorm:"column:maximum_spend; type:decimal(20,6)"`
}

func (Coupons) TableName() string {
	return "store_coupons"
}
