package prodvarsvc

import (
	"OnlineStoreBackend/models"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
)

func (service *Service) Delete(variationID uint64) {
	service.DB.Where("id = ?", variationID).Delete(&models.Variations{})
	detService := prodvardetsvc.NewServiceVariationDetail(service.DB)
	detService.Delete(variationID)
}
