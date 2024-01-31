package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseShippingData struct {
	ID             uint64  `json:"id"`
	ProductID      uint64  `json:"product_id"`
	Weight         float64 `json:"weight"`
	Dimension      string  `json:"dimension"`
	Classification string  `json:"classification"`
}

func NewResponseShippingData(c echo.Context, statusCode int, modelShipData models.ShippingData) error {
	responseShipData := ResponseShippingData{
		ID:             uint64(modelShipData.ID),
		Weight:         modelShipData.Weight,
		ProductID:      modelShipData.ProductID,
		Dimension:      modelShipData.Dimension,
		Classification: modelShipData.Classification,
	}
	return Response(c, statusCode, responseShipData)
}
