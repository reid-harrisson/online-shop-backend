package prodvardetsvc

import (
	"OnlineStoreBackend/models"
	"fmt"
)

func (service *Service) Create(variationID uint64, attributeValueIDs []uint64) {
	for _, attributeValueID := range attributeValueIDs {
		service.DB.Create(&models.ProductVariationDetails{
			VariationID:      variationID,
			AttributeValueID: attributeValueID,
		})
	}
}

func (service *Service) CreateWithCSV(modelNewDets *[]models.ProductVariationDetails, detMatches []string, detIndices map[string]int) {
	modelCurDets := []models.ProductVariationDetails{}
	service.DB.Where("Concat(variation_id,':',attribute_value_id) In (?)", detMatches).Find(&modelCurDets)
	for _, modelDet := range modelCurDets {
		match := fmt.Sprintf("%d:%d", modelDet.VariationID, modelDet.AttributeValueID)
		index := detIndices[match]
		(*modelNewDets)[index].ID = modelDet.ID
	}
	service.DB.Save(modelNewDets)
}
