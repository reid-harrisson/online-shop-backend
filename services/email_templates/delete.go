package etsvc

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

func (service *Service) Delete(templateID uint64) error {
	query := service.DB.Delete(&models.EmailTemplates{
		Model: gorm.Model{
			ID: uint(templateID),
		},
	})
	if query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
