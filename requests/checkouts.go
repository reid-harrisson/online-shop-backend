package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCheckout struct {
	BillingAddressID  uint64   `json:"billing_address_id" example:"1"`
	ShippingAddressID uint64   `json:"shipping_address_id" example:"1"`
	CouponIDs         []uint64 `json:"coupon_code" example:"123"`
	CardNumber        string   `json:"card_number" example:"4242424242424242"`
	ExpMonth          int64    `json:"exp_month" example:"8"`
	ExpYear           int64    `json:"exp_year" example:"26"`
	CVC               string   `json:"cvc" example:"437"`
}

func (request RequestCheckout) RequestCheckoutValidate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CardNumber, validation.Required),
		validation.Field(&request.ExpMonth, validation.Required),
		validation.Field(&request.ExpYear, validation.Required),
		validation.Field(&request.CVC, validation.Required),
	)
}
