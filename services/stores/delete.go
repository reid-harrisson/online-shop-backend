package storesvc

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

func (service *Service) Delete(storeID uint64) error {
	query := service.DB.Delete(&models.Stores{
		Model: gorm.Model{
			ID: uint(storeID),
		},
	})
	if query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
