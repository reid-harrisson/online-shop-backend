package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestCategory struct {
	Name     string `json:"name" example:"pear"`
	ParentID uint64 `json:"parent_id" example:"2"`
}

type RequestProductCategory struct {
	CategoryIDs []uint64 `json:"category_ids" example:"1,2,3"`
}

func (request RequestCategory) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.ParentID, validation.Required),
	)
}
