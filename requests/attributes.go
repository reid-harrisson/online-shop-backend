package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestAttributeItem struct {
	Name string `json:"name" example:"size"`
	Unit string `json:"unit" example:""`
}

type RequestAttribute struct {
	Attributes []RequestAttributeItem `json:"attributes"`
}

func (request RequestAttributeItem) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
	)
}
