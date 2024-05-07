package prodattrvalsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Update(attributeID uint64, req *requests.RequestAttributeValue) {
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
	service.DB.Where("Concat(attribute_id, ':', attribute_value) In (?)", valMatches).Find(&modelCurVals)
	service.DB.Where("Concat(attribute_id, ':', attribute_value) Not In (?) And attribute_id = ?", valMatches, attributeID).Delete(&models.AttributeValues{})
	for _, modelVal := range modelCurVals {
		match := fmt.Sprintf("%d:%s", modelVal.AttributeID, modelVal.AttributeValue)
		index := valIndices[match]
		modelNewVals[index].ID = modelVal.ID
	}
	service.DB.Save(&modelNewVals)
}

func (service *Service) UpdateByID(attributeValueID uint64, value string) error {
	return service.DB.Model(&models.AttributeValues{}).
		Where("attribute_id = ?", attributeValueID).
		Update("attribute_value", value).Error
}
