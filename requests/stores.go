package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestStore struct {
	CompanyID            uint64 `json:"company_id" example:"2"`
	OwnerID              uint64 `json:"owner_id" example:"1427"`
	ShowStockLevelStatus int8   `json:"show_stock_level_status" example:"0"`
	ShowOutOfStockStatus int8   `json:"show_out_of_stock_status" example:"0"`
	IsBackOrder          int8   `json:"is_back_order" example:"0"`
	DeliveryPolicy       string `json:"delivery_policy" example:"example delivery policy"`
	ReturnsPolicy        string `json:"returns_policy" example:"example return policy"`
	Terms                string `json:"terms" example:"example terms"`
	Active               int8   `json:"active" example:"1"`
}

func (request RequestStore) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CompanyID, validation.Required),
		validation.Field(&request.OwnerID, validation.Required),
	)
}
