package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductTag struct {
	Tags []string `json:"tags"`
}

type ResponseTag struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewResponseProductTags(c echo.Context, statusCode int, modelTags []models.ProductTagsWithName) error {
	tags := make([]string, 0)
	for _, modelTag := range modelTags {
		tags = append(tags, modelTag.TagName)
	}
	return Response(c, statusCode, ResponseProductTag{Tags: tags})
}

func NewResponseTags(c echo.Context, statusCode int, modelTags []models.Tags) error {
	responseTags := []ResponseTag{}
	for _, modelTag := range modelTags {
		responseTags = append(responseTags, ResponseTag{
			ID:   uint64(modelTag.ID),
			Name: modelTag.Name,
		})
	}
	return Response(c, statusCode, responseTags)
}

func NewResponseTag(c echo.Context, statusCode int, modelTag models.Tags) error {
	return Response(c, statusCode, ResponseTag{
		ID:   uint64(modelTag.ID),
		Name: modelTag.Name,
	})
}
