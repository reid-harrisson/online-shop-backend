package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductVariation struct {
	ProductID     uint64   `json:"product_id"`
	AttributeID   uint64   `json:"attribute_id"`
	AttributeName string   `json:"attribute_name"`
	Variants      []string `json:"variants"`
}

func NewResponseProductVariations(c echo.Context, statusCode int, modelVars []models.ProductVariationsWithName) error {
	responseAttrs := make([]ResponseProductVariation, 0)
	vars := make(map[string][]string)
	for _, modelVar := range modelVars {
		vars[modelVar.AttributeName] = append(vars[modelVar.AttributeName], modelVar.Variant+modelVar.AttributeUnit)
	}
	tempID := uint64(0)
	for _, modelVar := range modelVars {
		if tempID != modelVar.AttributeID {
			responseAttrs = append(responseAttrs, ResponseProductVariation{
				ProductID:     modelVar.ProductID,
				AttributeID:   modelVar.AttributeID,
				AttributeName: modelVar.AttributeName,
				Variants:      vars[modelVar.AttributeName],
			})
			tempID = modelVar.AttributeID
		}
	}
	return Response(c, statusCode, responseAttrs)
}
