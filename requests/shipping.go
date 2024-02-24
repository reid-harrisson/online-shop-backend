package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestShippingOption struct {
	Method string `json:"method" example:"flat rate"`
}

type RequestShippingData struct {
	Weight         float64 `json:"weight" example:"1.35"`
	Width          float64 `json:"width" example:"58"`
	Height         float64 `json:"height" example:"118"`
	Depth          float64 `json:"depth" example:"8"`
	Classification string  `json:"classification" example:"food"`
}

func (request RequestShippingOption) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Method, validation.Required),
	)
}

func (request RequestShippingData) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Weight, validation.Required),
		validation.Field(&request.Width, validation.Required),
		validation.Field(&request.Height, validation.Required),
		validation.Field(&request.Depth, validation.Required),
		validation.Field(&request.Classification, validation.Required),
	)
}
