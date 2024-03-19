package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestShippingData struct {
	Weight float64 `json:"weight" example:"1.35"`
	Width  float64 `json:"width" example:"58"`
	Height float64 `json:"height" example:"118"`
	Length float64 `json:"length" example:"8"`
}

func (request RequestShippingData) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Weight, validation.Required),
		validation.Field(&request.Width, validation.Required),
		validation.Field(&request.Height, validation.Required),
		validation.Field(&request.Length, validation.Required),
	)
}
