package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseShippingData struct {
	ID     uint64  `json:"id"`
	Weight float64 `json:"weight"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
}

type ResponseProductShippingData struct {
	ID        uint64  `json:"id"`
	ProductID uint64  `json:"product_id"`
	Weight    float64 `json:"weight"`
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	Length    float64 `json:"length"`
}

func NewResponseShippingData(c echo.Context, statusCode int, modelShipData models.ShippingData) error {
	responseShipData := ResponseProductShippingData{
		ID:        uint64(modelShipData.ID),
		Weight:    modelShipData.Weight,
		ProductID: modelShipData.VariationID,
		Width:     modelShipData.Width,
		Height:    modelShipData.Height,
		Length:    modelShipData.Length,
	}
	return Response(c, statusCode, responseShipData)
}
