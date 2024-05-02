package catesvc

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

func (service *Service) Delete(categoryID uint64) error {
	err := service.DB.Where("category_id = ?", categoryID).Delete(&models.ProductCategories{}).Error
	if err != nil {
		return err
	}

	modelCategories := make([]models.Categories, 0)
	service.DB.Where("parent_id = ?", categoryID).Find(&modelCategories)
	for _, modelCategory := range modelCategories {
		service.Delete(uint64(modelCategory.ID))
	}

	query := service.DB.Delete(&models.Categories{
		Model: gorm.Model{
			ID: uint(categoryID),
		},
	})
	if query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
