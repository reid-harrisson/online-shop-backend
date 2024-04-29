package prodvardetsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Update(variationID uint64, attributeValueIDs []uint64) error {
	for _, attributeValueID := range attributeValueIDs {
		return service.DB.Create(&models.VariationDetails{
			VariationID:      variationID,
			AttributeValueID: attributeValueID,
		}).Error
	}

	return nil
}
