package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductAttributeValue struct {
	ProductID     uint64   `json:"product_id"`
	AttributeID   uint64   `json:"attribute_id"`
	AttributeName string   `json:"attribute_name"`
	Values        []string `json:"values"`
}

func NewResponseProductAttributeValue(c echo.Context, statusCode int, modelValues []models.ProductAttributeValuesWithDetail) error {
	responseValues := make([]ResponseProductAttributeValue, 0)
	mapValues := make(map[string][]string)
	mapIndexes := make(map[string]int)
	for index, modelValue := range modelValues {
		mapValues[modelValue.AttributeName] = append(mapValues[modelValue.AttributeName], modelValue.AttributeValue+modelValue.Unit)
		mapIndexes[modelValue.AttributeName] = index
	}
	for name, values := range mapValues {
		index := mapIndexes[name]
		responseValues = append(responseValues, ResponseProductAttributeValue{
			ProductID:     modelValues[index].ProductID,
			AttributeID:   modelValues[index].AttributeID,
			AttributeName: name,
			Values:        values,
		})
	}
	return Response(c, statusCode, responseValues)
}
