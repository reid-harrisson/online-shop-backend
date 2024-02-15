package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductAttribute struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Name      string `json:"name"`
	Unit      string `json:"unit"`
}

func NewResponseProductAttribute(c echo.Context, statusCode int, modelAttr models.ProductAttributes) error {
	return Response(c, statusCode, ResponseProductAttribute{
		ID:        uint64(modelAttr.ID),
		ProductID: modelAttr.ProductID,
		Name:      modelAttr.Name,
		Unit:      modelAttr.Unit,
	})
}

func NewResponseProductAttributes(c echo.Context, statusCode int, modelAttrs []models.ProductAttributes) error {
	responseAttrs := make([]ResponseProductAttribute, 0)
	for _, modelAttr := range modelAttrs {
		responseAttrs = append(responseAttrs, ResponseProductAttribute{
			ID:        uint64(modelAttr.ID),
			ProductID: modelAttr.ProductID,
			Name:      modelAttr.Name,
			Unit:      modelAttr.Unit,
		})
	}
	return Response(c, statusCode, responseAttrs)
}
