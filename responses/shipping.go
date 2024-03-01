package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

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

type ResponseShippingMethod struct {
	Method        string  `json:"method"`
	FlatRate      float64 `json:"flat_rate"`
	BaseRate      float64 `json:"base_rate"`
	RatePerItem   float64 `json:"rate_per_item"`
	RatePerWeight float64 `json:"rate_per_weight"`
	RatePerTotal  float64 `json:"rate_per_total"`
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

func NewResponseShippingMethod(c echo.Context, statusCode int, modelMethods []models.ShippingMethods) error {
	responseMethods := make([]ResponseShippingMethod, 0)
	for _, modelMethod := range modelMethods {
		responseMethods = append(responseMethods, ResponseShippingMethod{
			Method:        utils.ShippingMethodsToString(modelMethod.Method),
			FlatRate:      modelMethod.FlatRate,
			BaseRate:      modelMethod.BaseRate,
			RatePerItem:   modelMethod.RatePerItem,
			RatePerWeight: modelMethod.RatePerWeight,
			RatePerTotal:  modelMethod.RatePerTotal,
		})
	}
	return Response(c, statusCode, responseMethods)
}
