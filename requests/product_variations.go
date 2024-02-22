package requests

import "OnlineStoreBackend/pkgs/utils"

type RequestStoreProductVariation struct {
	AttributeValueIDs []uint64            `json:"attribute_value_ids" example:"1,2,3"`
	Price             float64             `json:"price" example:"1.23"`
	StockLevel        float64             `json:"stock_level" example:"30"`
	Discount          float64             `json:"discount" example:"10"`
	Type              utils.DiscountTypes `json:"type" example:"1"`
	FreeShipping      float64             `json:"free_shipping" example:"20"`
}
