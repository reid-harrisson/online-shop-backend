package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupons struct {
	gorm.Model

	StoreID           uint64    `gorm:"column:store_id"`
	CouponCode        string    `gorm:"column:coupon_code"`
	DiscountType      int8      `gorm:"column:discount_type"`
	CouponAmount      float64   `gorm:"column:coupon_amount"`
	AllowFreeShipping int8      `gorm:"column:allow_free_shipping"`
	ExpiryDate        time.Time `gorm:"column:expiry_date"`
	MinimumSpend      float64   `gorm:"column:minimum_spend"`
	MaximumSpend      float64   `gorm:"column:maximum_spend"`
}

func (Coupons) TableName() string {
	return "store_coupons"
}
