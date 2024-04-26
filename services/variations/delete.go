package prodvarsvc

import (
	"OnlineStoreBackend/models"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
)

func (service *Service) Delete(variationID uint64) error {
	if err := service.DB.Where("id = ?", variationID).Delete(&models.Variations{}).Error; err != nil {
		return err
	}

	detService := prodvardetsvc.NewServiceVariationDetail(service.DB)
	if err := detService.Delete(variationID); err != nil {
		return err
	}

	return nil
}
