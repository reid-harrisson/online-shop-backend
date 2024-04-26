package prodattrvalsvc

import (
	"OnlineStoreBackend/models"
	"fmt"
)

func (service *Service) Create(attributeID uint64, value string) error {
	return service.DB.Create(&models.AttributeValues{
		AttributeID:    attributeID,
		AttributeValue: value,
	}).Error
}

func (service *Service) CreateWithCSV(modelNewVals *[]models.AttributeValues, valMatches []string, valIndices map[string]int) error {
	modelCurVals := []models.AttributeValues{}
	if err := service.DB.Where("Concat(attribute_id, ':', attribute_value) In (?)", valMatches).Find(&modelCurVals).Error; err != nil {
		return err
	}
	for _, modelVal := range modelCurVals {
		match := fmt.Sprintf("%d:%s", modelVal.AttributeID, modelVal.AttributeValue)
		index := valIndices[match]
		(*modelNewVals)[index].ID = modelVal.ID
	}
	return service.DB.Save(modelNewVals).Error
}
