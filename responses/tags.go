package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseTag struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	TagID     uint64 `json:"tag_id"`
	TagName   string `json:"tag_name"`
}

func NewResponseProductTags(c echo.Context, statusCode int, modelTags []models.ProductTagsWithName) error {
	responseTags := make([]ResponseTag, 0)
	for _, modelTag := range modelTags {
		responseTags = append(responseTags, ResponseTag{
			ID:        uint64(modelTag.ID),
			ProductID: modelTag.ProductID,
			TagID:     modelTag.TagID,
			TagName:   modelTag.TagName,
		})
	}
	return Response(c, statusCode, responseTags)
}
