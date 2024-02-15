package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCategory struct {
	StoreID uint64 `json:"store_id" example:"1"`
	Name    string `json:"name" example:"clothes, shirt"`
}

type RequestProductCategory struct {
	CategoryIDs []uint64 `json:"category_ids" example:"1,2,3"`
}

func (request RequestCategory) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.StoreID, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)
}
