package etsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(templateID uint64) error {
	return service.DB.Where("id = ?", templateID).Delete(&models.EmailTemplates{}).Error
}
