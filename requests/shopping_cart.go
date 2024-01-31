package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCartItem struct {
	ProductID  uint64  `json:"product_id" example:"1"`
	CustomerID uint64  `json:"customer_id" example:"1"`
	Quantity   float64 `json:"quantity" example:"1.0"`
}

func (request RequestCartItem) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CustomerID, validation.Required),
		validation.Field(&request.ProductID, validation.Required),
		validation.Field(&request.Quantity, validation.Required),
	)
}
