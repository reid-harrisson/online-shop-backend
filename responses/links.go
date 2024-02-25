package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseLinkedProducts struct {
	Upsell    []ResponseProduct `json:"up_sell"`
	CrossSell []ResponseProduct `json:"cross_sell"`
}

func NewResponseLinkedProducts(c echo.Context, statusCode int, modelLinks []models.ProductsWithLink) error {
	upSell := make([]ResponseProduct, 0)
	crossSell := make([]ResponseProduct, 0)
	for _, modelItem := range modelLinks {
		imageUrls := make([]string, 0)
		json.Unmarshal([]byte(modelItem.ImageUrls), &imageUrls)
		responseProduct := ResponseProduct{
			ID:               uint64(modelItem.ID),
			StoreID:          modelItem.StoreID,
			Title:            modelItem.Title,
			ShortDescription: modelItem.ShortDescription,
			LongDescription:  modelItem.LongDescription,
			ImageUrls:        imageUrls,
			Status:           utils.ProductStatusToString(modelItem.Status),
		}
		if modelItem.IsUpCross == utils.UpSell {
			upSell = append(upSell, responseProduct)
		} else {
			crossSell = append(crossSell, responseProduct)
		}
	}
	return Response(c, statusCode, ResponseLinkedProducts{
		Upsell:    upSell,
		CrossSell: crossSell,
	})
}
