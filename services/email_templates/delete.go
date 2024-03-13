package etsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(storeID uint64, templateID uint64) error {
	return service.DB.Where("id = ? And store_id = ?", templateID, storeID).Delete(&models.EmailTemplates{}).Error
}
