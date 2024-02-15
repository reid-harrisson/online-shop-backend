package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProductRate struct {
	ProductID  uint64  `json:"product_id" example:"1"`
	CustomerID uint64  `json:"customer_id" example:"1"`
	Rate       float64 `json:"rate" example:"1.0"`
}

func (request RequestProductRate) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CustomerID, validation.Required),
		validation.Field(&request.ProductID, validation.Required),
		validation.Field(&request.Rate, validation.Required),
	)
}
