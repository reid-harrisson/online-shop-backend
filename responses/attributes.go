package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseAttribute struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Name      string `json:"name"`
}

func NewResponseAttribute(c echo.Context, statusCode int, modelAttr models.Attributes) error {
	return Response(c, statusCode, ResponseAttribute{
		ID:        uint64(modelAttr.ID),
		ProductID: modelAttr.ProductID,
		Name:      modelAttr.AttributeName,
	})
}

func NewResponseAttributes(c echo.Context, statusCode int, modelAttrs []models.Attributes) error {
	responseAttrs := make([]ResponseAttribute, 0)
	for _, modelAttr := range modelAttrs {
		responseAttrs = append(responseAttrs, ResponseAttribute{
			ID:        uint64(modelAttr.ID),
			ProductID: modelAttr.ProductID,
			Name:      modelAttr.AttributeName,
		})
	}
	return Response(c, statusCode, responseAttrs)
}
