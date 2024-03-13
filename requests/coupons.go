package requests

type RequestCoupon struct {
	CouponCode        string  `json:"coupon_code" example:"123"`
	DiscountType      string  `json:"discount_type" example:"1"`
	CouponAmount      float64 `json:"coupon_amount" example:"12"`
	AllowFreeShipping int8    `json:"allow_free_shipping" example:"1"`
	ExpiryDate        string  `json:"expiry_date" example:"2024-05-01"`
	MinimumSpend      float64 `json:"minimum_spend" example:"100"`
	MaximumSpend      float64 `json:"maximum_spend" example:"200"`
}
