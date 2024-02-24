package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestProductOrderStatus struct {
	Status string `json:"status" example:"Processing"`
}

func (request RequestProductOrderStatus) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Status, validation.Required),
	)
}
