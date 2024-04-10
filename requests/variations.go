package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestVariation struct {
	AttributeValueIDs []uint64            `json:"attribute_value_ids" example:"1,2,3"`
	Price             float64             `json:"price" example:"1.23"`
	StockLevel        float64             `json:"stock_level" example:"30"`
	DiscountAmount    float64             `json:"discount_amount" example:"10"`
	DiscountType      utils.DiscountTypes `json:"discount_type" example:"1"`
	ImageUrls         []string            `json:"image_urls" example:"image1,image2,image3"`
	Description       string              `json:"description" example:"example description"`
	BackOrderAllowed  int8                `json:"back_order_allowed" example:"1"`
}

func (request RequestVariation) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.StockLevel, validation.Required),
	)
}

type RequestAttributeValueIDs struct {
	ValueIDs []uint64 `json:"attribute_value_ids" example:"5"`
}
