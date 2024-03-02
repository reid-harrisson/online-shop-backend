package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestShippingMethod struct {
	Method        string  `json:"method" example:"FlatRate"`
	FlatRate      float64 `json:"flat_rate" example:"5"`
	BaseRate      float64 `json:"base_rate" example:"2"`
	RatePerItem   float64 `json:"rate_per_item" example:"0.2"`
	RatePerWeight float64 `json:"rate_per_weight" example:"0.1"`
	RatePerTotal  float64 `json:"rate_per_total" example:"5"`
}

type RequestShippingData struct {
	Weight         float64 `json:"weight" example:"1.35"`
	Width          float64 `json:"width" example:"58"`
	Height         float64 `json:"height" example:"118"`
	Length         float64 `json:"length" example:"8"`
	Classification string  `json:"classification" example:"food"`
}

func (request RequestShippingMethod) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Method, validation.Required),
	)
}

func (request RequestShippingData) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Weight, validation.Required),
		validation.Field(&request.Width, validation.Required),
		validation.Field(&request.Height, validation.Required),
		validation.Field(&request.Length, validation.Required),
		validation.Field(&request.Classification, validation.Required),
	)
}
