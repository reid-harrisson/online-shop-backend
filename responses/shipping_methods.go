package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseShippingLocalPickup struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ZoneID      uint64  `json:"zone_id"`
	StoreID     uint64  `json:"store_id"`
	Method      string  `json:"method"`
	TaxStatus   int8    `json:"tax_status"`
	Cost        float64 `json:"cost"`
}

type ResponseShippingFlatRate struct {
	ID          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ZoneID      uint64  `json:"zone_id"`
	StoreID     uint64  `json:"store_id"`
	Method      string  `json:"method"`
	TaxStatus   int8    `json:"tax_status"`
	Cost        float64 `json:"cost"`
	Rates       []ResponseFlatRate
}

type ResponseShippingFree struct {
	ID                 uint64  `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	ZoneID             uint64  `json:"zone_id"`
	StoreID            uint64  `json:"store_id"`
	Method             string  `json:"method"`
	Requirement        string  `json:"requirement"`
	MinimumOrderAmount float64 `json:"minimum_order_amount"`
}

func NewResponseShippingLocalPickup(c echo.Context, statusCode int, modelMethod models.ShippingMethods) error {
	return Response(c, statusCode, ResponseShippingLocalPickup{
		ID:          uint64(modelMethod.ID),
		Title:       modelMethod.Title,
		Description: modelMethod.Description,
		StoreID:     modelMethod.StoreID,
		ZoneID:      modelMethod.ZoneID,
		Method:      utils.ShippingMethodsToString(modelMethod.Method),
		TaxStatus:   modelMethod.TaxStatus,
		Cost:        modelMethod.Cost,
	})
}

func NewResponseShippingFree(c echo.Context, statusCode int, modelMethod models.ShippingMethods) error {
	return Response(c, statusCode, ResponseShippingFree{
		ID:                 uint64(modelMethod.ID),
		Title:              modelMethod.Title,
		Description:        modelMethod.Description,
		StoreID:            modelMethod.StoreID,
		ZoneID:             modelMethod.ZoneID,
		Method:             utils.ShippingMethodsToString(modelMethod.Method),
		Requirement:        utils.RequirementToString(modelMethod.Requirement),
		MinimumOrderAmount: modelMethod.MinimumOrderAmount,
	})
}

func NewResponseShippingFlatRate(c echo.Context, statusCode int, modelMethod models.ShippingMethods, modelRates []models.ShippingFlatRates, modelClasses []models.ShippingClasses) error {
	indices := map[uint64]int{}
	for index, modelClass := range modelClasses {
		indices[uint64(modelClass.ID)] = index
	}
	responseRates := []ResponseFlatRate{}
	for _, modelRate := range modelRates {
		responseRates = append(responseRates, ResponseFlatRate{
			ID:          uint64(modelRate.ID),
			ClassID:     modelRate.ClassID,
			CostPerItem: modelRate.CostPerItem,
			BaseCost:    modelRate.BaseCost,
			Percent:     modelRate.Percent,
			MinFee:      modelRate.MinFee,
			MaxFee:      modelRate.MaxFee,
			ClassName:   modelClasses[indices[modelRate.ClassID]].Name,
		})
	}
	return Response(c, statusCode, ResponseShippingFlatRate{
		ID:          uint64(modelMethod.ID),
		Title:       modelMethod.Title,
		Description: modelMethod.Description,
		StoreID:     modelMethod.StoreID,
		ZoneID:      modelMethod.ZoneID,
		Method:      utils.ShippingMethodsToString(modelMethod.Method),
		TaxStatus:   modelMethod.TaxStatus,
		Cost:        modelMethod.Cost,
		Rates:       responseRates,
	})
}
