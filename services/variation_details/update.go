package prodvardetsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Update(variationID uint64, attributeValueIDs []uint64) error {
	modelDets := []models.VariationDetails{}
	modelCurDets := []models.VariationDetails{}
	detIndices := map[uint64]int{}

	err := service.DB.Where("attribute_value_id Not In (?) And variation_id = ?", attributeValueIDs, variationID).Delete(&models.VariationDetails{}).Error
	if err != nil {
		return err
	}

	err = service.DB.Where("variation_id = ?", variationID).Find(&modelCurDets).Error
	if err != nil {
		return err
	}

	for index, attributeValueID := range attributeValueIDs {
		detIndices[attributeValueID] = index + 1
		modelDets = append(modelDets, models.VariationDetails{
			VariationID:      variationID,
			AttributeValueID: attributeValueID,
		})
	}

	for _, modelDet := range modelCurDets {
		index := detIndices[modelDet.AttributeValueID]
		if index > 0 {
			modelDets[index-1].ID = modelDet.ID
		}
	}

	if len(modelDets) == 0 {
		return nil
	}

	return service.DB.Save(&modelDets).Error
}
