package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCoupons struct {
	StoreID uint64  `json:"store_id" example:"1"`
	Code    string  `json:"code" example:"20percent"`
	Rule    string  `json:"type" example:"percent"`
	Amount  float64 `json:"amount" example:"10"`
}

func (request RequestCoupons) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.StoreID, validation.Required),
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Rule, validation.Required),
		validation.Field(&request.Amount, validation.Required),
	)
}
