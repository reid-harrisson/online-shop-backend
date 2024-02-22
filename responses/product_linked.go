package responses

import (
	"OnlineStoreBackend/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseProductLinked struct {
	Upsell    []ResponseProduct `json:"up_sell"`
	CrossSell []ResponseProduct `json:"cross_sell"`
}

func NewResponseProductLinked(c echo.Context, statusCode int, modelProductLinked []models.ProductsWithLink) error {
	upSell := make([]ResponseProduct, 0)
	crossSell := make([]ResponseProduct, 0)
	for _, modelItem := range modelProductLinked {
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelItem.ImageUrls), &imageUrls)
		responseProduct := ResponseProduct{
			ID:               uint64(modelItem.ID),
			StoreID:          modelItem.StoreID,
			Title:            modelItem.Title,
			ShortDescription: modelItem.ShortDescription,
			LongDescription:  modelItem.LongDescription,
			SKU:              modelItem.SKU,
			ImageUrls:        imageUrls,
			UnitPriceRegular: modelItem.UnitPriceRegular,
			UnitPriceSale:    modelItem.UnitPriceSale,
			StockQuantity:    modelItem.StockQuantity,
			Active:           modelItem.Active,
		}
		if modelItem.IsUpCross == models.UpSell {
			upSell = append(upSell, responseProduct)
		} else {
			crossSell = append(crossSell, responseProduct)
		}
	}
	return Response(c, statusCode, ResponseProductLinked{
		Upsell:    upSell,
		CrossSell: crossSell,
	})
}
