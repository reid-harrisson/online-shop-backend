package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseStoreProductVariation struct {
	ID         uint64  `json:"id"`
	Sku        string  `json:"sku"`
	ProductID  uint64  `json:"product_id"`
	Price      float64 `json:"price"`
	StockLevel float64 `json:"stock_level"`
}

func NewResponseStoreProductVariation(c echo.Context, statusCode int, modelVar models.StoreProductVariations) error {
	return Response(c, statusCode, ResponseStoreProductVariation{
		ID:         uint64(modelVar.ID),
		Sku:        modelVar.Sku,
		ProductID:  modelVar.ProductID,
		Price:      modelVar.Price,
		StockLevel: modelVar.StockLevel,
	})
}
