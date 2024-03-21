package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestStockTrack struct {
	ProductID   uint64            `json:"product_id"`
	VariationID uint64            `json:"variation_id"`
	Change      float64           `json:"change"`
	Event       utils.TrackEvents `json:"event"`
}

func (request RequestStockTrack) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ProductID, validation.Required),
		validation.Field(&request.VariationID, validation.Required),
	)
}
