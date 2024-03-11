package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseTag struct {
	Tags []string `json:"tags"`
}

type ResponseStoreTag struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewResponseProductTags(c echo.Context, statusCode int, modelTags []models.ProductTagsWithName) error {
	tags := make([]string, 0)
	for _, modelTag := range modelTags {
		tags = append(tags, modelTag.TagName)
	}
	return Response(c, statusCode, ResponseTag{Tags: tags})
}

func NewResponseStoreTags(c echo.Context, statusCode int, modelTags []models.StoreTags) error {
	responseTags := []ResponseStoreTag{}
	for _, modelTag := range modelTags {
		responseTags = append(responseTags, ResponseStoreTag{
			ID:   uint64(modelTag.ID),
			Name: modelTag.Name,
		})
	}
	return Response(c, statusCode, responseTags)
}
