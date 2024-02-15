package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductAttribute struct {
	ID            uint64 `json:"id"`
	ProductID     uint64 `json:"product_id"`
	AttributeID   uint64 `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
}

func NewResponseProductAttributes(c echo.Context, statusCode int, modelAttrs []models.ProductAttributesWithName) error {
	responseAttrs := make([]ResponseProductAttribute, 0)
	for _, modelAttr := range modelAttrs {
		responseAttrs = append(responseAttrs, ResponseProductAttribute{
			ID:            uint64(modelAttr.ID),
			ProductID:     modelAttr.ProductID,
			AttributeID:   modelAttr.AttributeID,
			AttributeName: modelAttr.AttributeName,
		})
	}
	return Response(c, statusCode, responseAttrs)
}
