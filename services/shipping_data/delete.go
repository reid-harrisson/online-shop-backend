package shipsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(variationID uint64) error {
	return service.DB.Where("variation_id = ?", variationID).Delete(&models.ShippingData{}).Error
}
