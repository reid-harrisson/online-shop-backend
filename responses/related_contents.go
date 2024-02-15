package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductContent struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	ContentID uint64 `json:"content_id"`
}

func NewResponseProductContents(c echo.Context, statusCode int, modelContents []models.ProductContents) error {
	responseContents := make([]ResponseProductContent, 0)
	for _, modelContent := range modelContents {
		responseContents = append(responseContents, ResponseProductContent{
			ID:        uint64(modelContent.ID),
			ProductID: modelContent.ProductID,
			ContentID: modelContent.ContentID,
		})
	}
	return Response(c, statusCode, responseContents)
}
