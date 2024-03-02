package prodcatesvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(categoryID uint64, productID uint64) {
	service.DB.Create(&models.ProductCategories{
		CategoryID: categoryID,
		ProductID:  productID,
	})
}

func (service *Service) CreateWithCSV(modelCategories []models.StoreCategories, productID uint64) {
	for _, modelCategory := range modelCategories {
		service.DB.Create(&models.ProductCategories{
			CategoryID: uint64(modelCategory.ID),
			ProductID:  productID,
		})
	}
}
