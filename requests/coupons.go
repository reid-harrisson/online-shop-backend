package requests

import (
	"time"
)

type RequestCoupon struct {
	CouponCode        string    `json:"coupon_code"`
	DiscountType      int8      `json:"discount_type"`
	CouponAmount      float64   `json:"coupon_amount"`
	AllowFreeShipping int8      `json:"allow_free_shipping"`
	ExpiryDate        time.Time `json:"expiry_date"`
	MinimumSpend      float64   `json:"minimum_spend"`
	MaximumSpend      float64   `json:"maximum_spend"`
}
