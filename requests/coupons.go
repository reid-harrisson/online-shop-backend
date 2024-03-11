package requests

type RequestCoupon struct {
	CouponCode        string  `json:"coupon_code" example:"123"`
	DiscountType      int8    `json:"discount_type" example:"1"`
	CouponAmount      float64 `json:"coupon_amount" example:"12"`
	AllowFreeShipping int8    `json:"allow_free_shipping" example:"1"`
	ExpiryDate        string  `json:"expiry_date" example:"2003-3-4"`
	MinimumSpend      float64 `json:"minimum_spend" example:"1"`
	MaximumSpend      float64 `json:"maximum_spend" example:"1"`
}
