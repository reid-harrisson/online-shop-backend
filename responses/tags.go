package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseTag struct {
	Tags []string `json:"tags"`
}

func NewResponseProductTags(c echo.Context, statusCode int, modelTags []models.ProductTagsWithName) error {
	tags := make([]string, 0)
	for _, modelTag := range modelTags {
		tags = append(tags, modelTag.TagName)
	}
	return Response(c, statusCode, ResponseTag{Tags: tags})
}
