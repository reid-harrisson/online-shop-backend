package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseShippingMethod struct {
	Method              string  `json:"method"`
	Requirement         string  `json:"requirement"`
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

type ResponseShippingLocalPickup struct {
	ID        uint64  `json:"id"`
	ZoneID    uint64  `json:"zone_id"`
	StoreID   uint64  `json:"store_id"`
	Method    string  `json:"method"`
	TaxStatus int8    `json:"tax_status"`
	Cost      float64 `json:"cost"`
}

type ResponseShippingFree struct {
	ID                 uint64  `json:"id"`
	ZoneID             uint64  `json:"zone_id"`
	StoreID            uint64  `json:"store_id"`
	Method             string  `json:"method"`
	Requirement        string  `json:"requirement"`
	MinimumOrderAmount float64 `json:"minimum_order_amount"`
}

func NewResponseShippingMethod(c echo.Context, statusCode int, modelMethods []models.ShippingMethods) error {
	responseMethods := make([]ResponseShippingMethod, 0)
	for _, modelMethod := range modelMethods {
		responseMethods = append(responseMethods, ResponseShippingMethod{
			Method:              utils.ShippingMethodsToString(modelMethod.Method),
			Requirement:         utils.RequirementToString(modelMethod.Requirement),
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

func NewResponseShippingLocalPickup(c echo.Context, statusCode int, modelMethod models.ShippingMethods) error {
	return Response(c, statusCode, ResponseShippingLocalPickup{
		ID:        uint64(modelMethod.ID),
		StoreID:   modelMethod.StoreID,
		ZoneID:    modelMethod.ZoneID,
		Method:    utils.ShippingMethodsToString(modelMethod.Method),
		TaxStatus: modelMethod.TaxStatus,
		Cost:      modelMethod.Cost,
	})
}

func NewResponseShippingFree(c echo.Context, statusCode int, modelMethod models.ShippingMethods) error {
	return Response(c, statusCode, ResponseShippingFree{
		ID:                 uint64(modelMethod.ID),
		StoreID:            modelMethod.StoreID,
		ZoneID:             modelMethod.ZoneID,
		Method:             utils.ShippingMethodsToString(modelMethod.Method),
		Requirement:        utils.RequirementToString(modelMethod.Requirement),
		MinimumOrderAmount: modelMethod.MinimumOrderAmount,
	})
}
