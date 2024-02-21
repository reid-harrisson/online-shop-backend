package etsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(emailTemplateID uint64) error {
	return service.DB.
		Where("id = ?", emailTemplateID).
		Delete(models.EmailTemplate{}).
		Error
}
