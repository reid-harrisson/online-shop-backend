package prodvardetsvc

import (
	"OnlineStoreBackend/models"
	"fmt"
)

func (service *Service) Create(variationID uint64, attributeValueIDs []uint64) {
	modelDets := []models.VariationDetails{}
	for _, attributeValueID := range attributeValueIDs {
		modelDets = append(modelDets, models.VariationDetails{
			VariationID:      variationID,
			AttributeValueID: attributeValueID,
		})
	}
	service.DB.Create(&modelDets)
}

func (service *Service) CreateWithCSV(modelNewDets *[]models.VariationDetails, detMatches []string, detIndices map[string]int) error {
	modelCurDets := []models.VariationDetails{}
	if err := service.DB.Where("Concat(variation_id,':',attribute_value_id) In (?)", detMatches).Find(&modelCurDets).Error; err != nil {
		return err
	}
	for _, modelDet := range modelCurDets {
		match := fmt.Sprintf("%d:%d", modelDet.VariationID, modelDet.AttributeValueID)
		index := detIndices[match]
		(*modelNewDets)[index].ID = modelDet.ID
	}
	return service.DB.Save(modelNewDets).Error
}
