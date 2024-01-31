package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductTag struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Tag       string `json:"tag"`
}

func NewResponseProductTags(c echo.Context, statusCode int, modelTags []models.Tags) error {
	responseTags := make([]ResponseProductTag, 0)
	for _, modelTag := range modelTags {
		responseTags = append(responseTags, ResponseProductTag{
			ID:        uint64(modelTag.ID),
			ProductID: modelTag.ProductID,
			Tag:       modelTag.Tag,
		})
	}
	return Response(c, statusCode, responseTags)
}
