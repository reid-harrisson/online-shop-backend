package tagsvc

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

func (service *Service) Delete(tagID uint64) error {
	query := service.DB.Delete(&models.Tags{
		Model: gorm.Model{
			ID: uint(tagID),
		},
	})
	if query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
