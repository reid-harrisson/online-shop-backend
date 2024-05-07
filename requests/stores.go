package requests

import (
	"OnlineStoreBackend/pkgs/utils"
)

type RequestStore struct {
	Name                 string               `json:"name" example:"The Sample Shop"`
	ContactPhone         string               `json:"contact_phone" example:"+1234567890"`
	ContactEmail         string               `json:"contact_email" example:"example@sample.com"`
	ShowStockLevelStatus utils.SimpleStatuses `json:"show_stock_level_status" example:"0"`
	ShowOutOfStockStatus utils.SimpleStatuses `json:"show_out_of_stock_status" example:"0"`
	BackOrderStatus      utils.SimpleStatuses `json:"back_order_status" example:"0"`
	DeliveryPolicy       string               `json:"delivery_policy" example:"example delivery policy"`
	ReturnsPolicy        string               `json:"returns_policy" example:"example return policy"`
	Terms                string               `json:"terms" example:"example terms"`
}
