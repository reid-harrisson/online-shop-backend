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
	Active               int8                 `gorm:"column:active; example:0"`
	BackgroundColor1     string               `gorm:"column:background_color_1; example:#3b3939"`
	BackgroundColor2     string               `gorm:"column:background_color_2; example:#F2F2EE"`
	StoreBackground      string               `gorm:"column:store_background; example:/bg/background.png"`
	StoreLogo            string               `gorm:"column:store_logo; example:/logo/logo.png"`
	Description          string               `gorm:"column:description; example:This is example store for test."`
	HeaderLayoutStyle    int8                 `gorm:"column:header_layout_style; example:1"`
	ShowStoreLogo        int8                 `gorm:"column:show_store_logo; example:2"`
	ShowStoreTitleText   int8                 `gorm:"column:show_store_title_text; example:1"`
	Website              string               `gorm:"column:website; example:https://google.com/"`
	WebsiteButtonColor   string               `gorm:"column:website_button_color; example:#D3E2F1"`
}
