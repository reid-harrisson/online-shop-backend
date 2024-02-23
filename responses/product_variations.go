package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductVariation struct {
	ID         uint64  `json:"id"`
	Sku        string  `json:"sku"`
	ProductID  uint64  `json:"product_id"`
	Price      float64 `json:"price"`
	StockLevel float64 `json:"stock_level"`
}

type ResponseProductVariationsInStore struct {
	ProductID         uint64                     `json:"product_id"`
	Title             string                     `json:"title"`
	StockLevel        float64                    `json:"stock_level"`
	MinimumStockLevel float64                    `json:"minimum_stock_level"`
	Variations        []ResponseProductVariation `json:"variations"`
}

func NewResponseProductVariation(c echo.Context, statusCode int, modelVar models.ProductVariations) error {
	return Response(c, statusCode, ResponseProductVariation{
		ID:         uint64(modelVar.ID),
		Sku:        modelVar.Sku,
		ProductID:  modelVar.ProductID,
		Price:      modelVar.Price,
		StockLevel: modelVar.StockLevel,
	})
}

func NewResponseProductVariationsInStore(c echo.Context, statusCode int, modelVars []models.ProductVariationsWithDetail) error {
	mapVars := make(map[uint64][]int)
	responseProducts := make([]ResponseProductVariationsInStore, 0)
	for index, modelVar := range modelVars {
		mapVars[modelVar.ProductID] = append(mapVars[modelVar.ProductID], index)
	}
	for productID, indexes := range mapVars {
		responseVars := make([]ResponseProductVariation, 0)
		minimumStockLevel := float64(0)
		title := ""
		stockLevel := float64(0)
		for _, index := range indexes {
			if title == "" {
				minimumStockLevel = modelVars[index].MinimumStockLevel
				title = modelVars[index].Title
			}
			stockLevel += modelVars[index].StockLevel
			responseVars = append(responseVars, ResponseProductVariation{
				ID:         uint64(modelVars[index].ID),
				Sku:        modelVars[index].Sku,
				ProductID:  modelVars[index].ProductID,
				Price:      modelVars[index].Price,
				StockLevel: modelVars[index].StockLevel,
			})
		}
		responseProducts = append(responseProducts, ResponseProductVariationsInStore{
			MinimumStockLevel: minimumStockLevel,
			Title:             title,
			StockLevel:        stockLevel,
			ProductID:         productID,
			Variations:        responseVars,
		})
	}
	return Response(c, statusCode, responseProducts)
}
