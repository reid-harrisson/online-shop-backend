package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestShippingMethod struct {
	ProductID uint64 `json:"product_id" example:"1"`
	Method    string `json:"method" example:"flat"`
}

func (request RequestShippingMethod) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ProductID, validation.Required),
		validation.Field(&request.Method, validation.Required),
	)
}
