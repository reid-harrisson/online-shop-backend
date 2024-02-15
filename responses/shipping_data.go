package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseShippingData struct {
	ID             uint64  `json:"id"`
	Weight         float64 `json:"weight"`
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	Depth          float64 `json:"depth"`
	Classification string  `json:"classification"`
}

type ResponseProductShippingData struct {
	ID             uint64  `json:"id"`
	ProductID      uint64  `json:"product_id"`
	Weight         float64 `json:"weight"`
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	Depth          float64 `json:"depth"`
	Classification string  `json:"classification"`
}

func NewResponseShippingData(c echo.Context, statusCode int, modelShipData models.ShippingData) error {
	responseShipData := ResponseProductShippingData{
		ID:             uint64(modelShipData.ID),
		Weight:         modelShipData.Weight,
		ProductID:      modelShipData.ProductID,
		Width:          modelShipData.Width,
		Height:         modelShipData.Height,
		Depth:          modelShipData.Depth,
		Classification: modelShipData.Classification,
	}
	return Response(c, statusCode, responseShipData)
}
