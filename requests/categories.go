package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCategory struct {
	StoreID uint64 `json:"store_id" example:"1"`
	Name    string `json:"name" example:"clothes, shirt"`
}

func (request RequestCategory) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.StoreID, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)
}
