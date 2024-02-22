package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestStore struct {
	CompanyID        uint64                 `json:"company_id" example:"2"`
	OwnerID          uint64                 `json:"owner_id" example:"1427"`
	ContactPhone     string                 `json:"contact_phone" example:"7184756027"`
	ContactEmail     string                 `json:"contact_email" example:"example@sample.com"`
	StockLevelStatus utils.StockLevelStatus `json:"stock_level_status" example:"0"`
	OutOfStockStatus utils.OutOfStockStatus `json:"out_of_stock_status" example:"0"`
	BackOrderStatus  utils.BackOrderStatus  `json:"back_order_status" example:"0"`
	DeliveryPolicy   string                 `json:"delivery_policy" example:"example delivery policy"`
	ReturnsPolicy    string                 `json:"returns_policy" example:"example return policy"`
	Terms            string                 `json:"terms" example:"example terms"`
}

func (request RequestStore) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CompanyID, validation.Required),
		validation.Field(&request.OwnerID, validation.Required),
	)
}
