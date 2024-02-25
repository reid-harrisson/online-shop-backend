package storesvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(storeID uint64) error {
	return service.DB.Where("id = ?", storeID).Delete(&models.Stores{}).Error
}
