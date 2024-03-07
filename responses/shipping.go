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
	Method              string  `json:"method"`
	Requirement         int8    `json:"requirement"`
	MinimumOrderAmount  float64 `json:"minimum_order_amount"`
	TaxStatus           int8    `json:"tax_status"`
	Cost                float64 `json:"cost"`
	TaxIncluded         int8    `json:"tax_included"`
	HandlingFee         float64 `json:"handling_fee"`
	MaximumShippingCost float64 `json:"maximum_shipping_cost"`
	CalculationType     int8    `json:"calculation_type"`
	HandlingFeePerClass float64 `json:"handling_fee_per_class"`
	MinimumCostPerClass float64 `json:"minimum_cost_per_class"`
	MaximumCostPerClass float64 `json:"maximum_cost_per_class"`
	DiscountInMinMax    int8    `json:"discount_in_min_max"`
	TaxInMinMax         int8    `json:"tax_in_min_max"`
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
			Method:              utils.ShippingMethodsToString(modelMethod.Method),
			Requirement:         modelMethod.Requirement,
			MinimumOrderAmount:  modelMethod.MinimumOrderAmount,
			TaxStatus:           modelMethod.TaxStatus,
			Cost:                modelMethod.Cost,
			TaxIncluded:         modelMethod.TaxIncluded,
			HandlingFee:         modelMethod.HandlingFee,
			MaximumShippingCost: modelMethod.MaximumShippingCost,
			CalculationType:     modelMethod.CalculationType,
			HandlingFeePerClass: modelMethod.HandlingFeePerClass,
			MinimumCostPerClass: modelMethod.MinimumCostPerClass,
			MaximumCostPerClass: modelMethod.MaximumCostPerClass,
			DiscountInMinMax:    modelMethod.DiscountInMinMax,
			TaxInMinMax:         modelMethod.TaxInMinMax,
		})
	}
	return Response(c, statusCode, responseMethods)
}
