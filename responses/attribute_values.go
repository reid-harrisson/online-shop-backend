package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseAttributeValueItem struct {
	AttributeValueID uint64 `json:"attribute_value_id"`
	Value            string `json:"value"`
}

type ResponseAttributeValue struct {
	AttributeID   uint64                       `json:"attribute_id"`
	AttributeName string                       `json:"attribute_name"`
	Values        []ResponseAttributeValueItem `json:"values"`
}

func NewResponseAttributeValueByProduct(c echo.Context, statusCode int, modelValues []models.AttributeValuesWithDetail) error {
	responseValues := make([]ResponseAttributeValue, 0)
	mapValues := make(map[string][]ResponseAttributeValueItem)
	mapIndexes := make(map[string]int)
	for index, modelValue := range modelValues {
		mapValues[modelValue.AttributeName] = append(mapValues[modelValue.AttributeName], ResponseAttributeValueItem{
			AttributeValueID: uint64(modelValue.ID),
			Value:            modelValue.AttributeValue,
		})
		mapIndexes[modelValue.AttributeName] = index
	}
	for _, modelValue := range modelValues {
		if mapIndexes[modelValue.AttributeName] != -1 {
			responseValues = append(responseValues, ResponseAttributeValue{
				AttributeID:   modelValue.AttributeID,
				AttributeName: modelValue.AttributeName,
				Values:        mapValues[modelValue.AttributeName],
			})
			mapIndexes[modelValue.AttributeName] = -1
		}
	}
	return Response(c, statusCode, responseValues)
}
