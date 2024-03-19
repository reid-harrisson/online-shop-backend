package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseInventoryVariation struct {
	VariationID   uint64  `json:"variation_id"`
	VariationName string  `json:"variation_name"`
	StockLevel    float64 `json:"stock_level"`
}

type ResponseInventoryProduct struct {
	ProductID         uint64                       `json:"product_id"`
	ProductName       string                       `json:"product_name"`
	StockLevel        float64                      `json:"stock_level"`
	MinimumStockLevel float64                      `json:"minimum_stock_level"`
	Variations        []ResponseInventoryVariation `json:"variations"`
}

type ResponseInventory struct {
	Products   []ResponseInventoryProduct `json:"products"`
	StockLevel float64                    `json:"stock_level"`
}

func NewResponseInventory(c echo.Context, statusCode int, modelInvs []models.Inventories) error {
	prodIndices := map[uint64]int{}
	varIndices := map[uint64][]int{}
	responseProds := []ResponseInventoryProduct{}
	responseVars := []ResponseInventoryVariation{}
	for index, modelInv := range modelInvs {
		if prodIndices[modelInv.ProductID] == 0 {
			responseProds = append(responseProds, ResponseInventoryProduct{
				ProductID:         modelInv.ProductID,
				ProductName:       modelInv.ProductName,
				MinimumStockLevel: modelInv.MinimumStockLevel,
				StockLevel:        0,
				Variations:        []ResponseInventoryVariation{},
			})
			prodIndices[modelInv.ProductID] = len(responseProds)
		}
		varIndices[modelInv.ProductID] = append(varIndices[modelInv.ProductID], index)
		responseVars = append(responseVars, ResponseInventoryVariation{
			VariationID:   modelInv.VariationID,
			VariationName: modelInv.VariationName,
			StockLevel:    modelInv.StockLevel,
		})
	}
	stockLevel := 0.0
	for productID, indices := range varIndices {
		for _, index := range indices {
			prodIndex := prodIndices[productID] - 1
			responseProds[prodIndex].Variations = append(responseProds[prodIndex].Variations, responseVars[index])
			responseProds[prodIndex].StockLevel += responseVars[index].StockLevel
			stockLevel += responseVars[index].StockLevel
		}
	}

	return Response(c, statusCode, ResponseInventory{
		Products:   responseProds,
		StockLevel: stockLevel,
	})
}
