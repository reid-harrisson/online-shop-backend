package prodattrvalsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Update(attributeID uint64, req *requests.RequestAttributeValue) error {
	modelNewVals := []models.AttributeValues{}
	modelCurVals := []models.AttributeValues{}
	valIndices := map[string]int{}
	valMatches := []string{}
	for index, val := range req.Values {
		match := fmt.Sprintf("%d:%s", attributeID, val)
		valMatches = append(valMatches, match)
		valIndices[match] = index
		modelNewVals = append(modelNewVals, models.AttributeValues{
			AttributeID:    attributeID,
			AttributeValue: val,
		})
	}

	err := service.DB.Where("Concat(attribute_id, ':', attribute_value) In (?)", valMatches).Find(&modelCurVals).Error
	if err != nil {
		return err
	}

	err = service.DB.Where("Concat(attribute_id, ':', attribute_value) Not In (?) And attribute_id = ?", valMatches, attributeID).Delete(&models.AttributeValues{}).Error
	if err != nil {
		return err
	}

	for _, modelVal := range modelCurVals {
		match := fmt.Sprintf("%d:%s", modelVal.AttributeID, modelVal.AttributeValue)
		index := valIndices[match]
		modelNewVals[index].ID = modelVal.ID
	}

	return service.DB.Save(&modelNewVals).Error
}

func (service *Service) UpdateByID(attributeValueID uint64, value string) error {
	return service.DB.Model(&models.AttributeValues{}).
		Where("id = ?", attributeValueID).
		Update("attribute_value", value).Error
}
