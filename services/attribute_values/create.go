package prodattrvalsvc

import (
	"OnlineStoreBackend/models"
	"fmt"
)

func (service *Service) Create(attributeID uint64, value string) error {
	return service.DB.Create(&models.ProductAttributeValues{
		AttributeID:    attributeID,
		AttributeValue: value,
	}).Error
}

func (service *Service) CreateWithCSV(modelNewVals *[]models.ProductAttributeValues, valMatches []string, valIndices map[string]int) {
	modelCurVals := []models.ProductAttributeValues{}
	service.DB.Where("Concat(attribute_id, ':', attribute_value) In (?)", valMatches).Find(&modelCurVals)
	for _, modelVal := range modelCurVals {
		match := fmt.Sprintf("%d:%s", modelVal.AttributeID, modelVal.AttributeValue)
		index := valIndices[match]
		(*modelNewVals)[index].ID = modelVal.ID
	}
	service.DB.Save(modelNewVals)
}
