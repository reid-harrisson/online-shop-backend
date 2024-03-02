package prodvardetsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Update(variationID uint64, attributeValueIDs []uint64) {
	for _, attributeValueID := range attributeValueIDs {
		service.DB.Create(&models.ProductVariationDetails{
			VariationID:      variationID,
			AttributeValueID: attributeValueID,
		})
	}
}
