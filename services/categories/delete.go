package catesvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(categoryID uint64) {
	service.DB.Model(models.Products{}).Where("category_id = ?", categoryID).Update("category_id", 0)
	modelCategories := make([]models.BaseCategories, 0)
	service.DB.Where("parent_id = ?", categoryID).Find(&modelCategories)
	for _, modelCategory := range modelCategories {
		service.Delete(uint64(modelCategory.ID))
	}
	service.DB.Where("id = ?", categoryID).Delete(models.BaseCategories{})
}
