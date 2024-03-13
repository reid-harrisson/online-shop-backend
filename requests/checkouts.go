package requests

type RequestCheckout struct {
	BillingAddressID  uint64   `json:"billing_address_id" example:"1"`
	ShippingAddressID uint64   `json:"shipping_address_id" example:"1"`
	CouponIDs         []uint64 `json:"coupon_code" example:"123"`
	CardNumber        string   `json:"card_number" example:"4242424242424242"`
	ExpMonth          int64    `json:"exp_month" example:"8"`
	ExpYear           int64    `json:"exp_year" example:"26"`
	CVC               string   `json:"cvc" example:"437"`
}
