package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProductReview struct {
	Rate    float64 `json:"rate" example:"4.3"`
	Comment string  `json:"comment" example:"This is sample comment."`
}

type RequestProductReviewStatus struct {
	Status string `json:"status" example:"Published"`
}

func (request RequestProductReview) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Rate, validation.Required),
		validation.Field(&request.Comment, validation.Required),
	)
}

func (request RequestProductReviewStatus) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Status, validation.Required),
	)
}
