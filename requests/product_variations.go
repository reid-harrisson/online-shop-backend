package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestProductVariation struct {
	AttributeValueIDs []uint64            `json:"attribute_value_ids" example:"1,2,3"`
	Price             float64             `json:"price" example:"1.23"`
	StockLevel        float64             `json:"stock_level" example:"30"`
	Discount          float64             `json:"discount" example:"10"`
	Type              utils.DiscountTypes `json:"type" example:"1"`
	FreeShipping      float64             `json:"free_shipping" example:"20"`
}

func (request RequestProductVariation) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.StockLevel, validation.Required),
	)
}
