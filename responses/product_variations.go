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

type ResponseProductVariationWithProduct struct {
	ProductID  uint64                     `json:"product_id"`
	Variations []ResponseProductVariation `json:"variations"`
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

func NewResponseProductVariationWithProduct(c echo.Context, statusCode int, modelVars []models.ProductVariations) error {
	mapVars := make(map[uint64][]*models.ProductVariations)
	responseProducts := make([]ResponseProductVariationWithProduct, 0)
	for _, modelVar := range modelVars {
		mapVars[modelVar.ProductID] = append(mapVars[modelVar.ProductID], &modelVar)
	}
	for productID, modelVars := range mapVars {
		responseVars := make([]ResponseProductVariation, 0)
		for _, modelVar := range modelVars {
			responseVars = append(responseVars, ResponseProductVariation{
				ID:         uint64(modelVar.ID),
				Sku:        modelVar.Sku,
				ProductID:  modelVar.ProductID,
				Price:      modelVar.Price,
				StockLevel: modelVar.StockLevel,
			})
		}
		responseProducts = append(responseProducts, ResponseProductVariationWithProduct{
			ProductID:  productID,
			Variations: responseVars,
		})
	}
	return Response(c, statusCode, responseProducts)
}
