package catesvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(categoryID uint64) {
	service.DB.Where("category_id = ?", categoryID).Delete(&models.ProductCategories{})
	modelCategories := make([]models.Categories, 0)
	service.DB.Where("parent_id = ?", categoryID).Find(&modelCategories)
	for _, modelCategory := range modelCategories {
		service.Delete(uint64(modelCategory.ID))
	}
	service.DB.Where("id = ?", categoryID).Delete(&models.Categories{})
}
