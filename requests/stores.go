package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestStore struct {
	CompanyID              uint64  `json:"company_id" example:"2"`
	OwnerID                uint64  `json:"owner_id" example:"1427"`
	ContactPhone           string  `json:"contact_phone" example:"082 444 0107"`
	ContactEmail           string  `json:"contact_email" example:"davekeetis@gmail.com"`
	ShowStockQuantity      int8    `json:"show_stock_quantity_status" example:"0"`
	ShowOutOfStockProducts int8    `json:"show_out_of_stock_status" example:"0"`
	DeliveryPolicy         string  `json:"delivery_policy" example:"example delivery policy"`
	ReturnsPolicy          string  `json:"returns_policy" example:"example return policy"`
	Terms                  string  `json:"terms" example:"example terms"`
	FlatRateShipping       float64 `json:"flat_rate_shipping" example:"0"`
	BackOrder              int8    `json:"back_order_status" example:"0"`
	Active                 int8    `json:"active" example:"1"`
}

func (request RequestStore) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.CompanyID, validation.Required),
		validation.Field(&request.OwnerID, validation.Required),
	)
}
