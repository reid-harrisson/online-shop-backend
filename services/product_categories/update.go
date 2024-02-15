package prodcatesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelCategories *[]models.ProductCategoriesWithName, req *requests.RequestProductCategory, productID uint64) {
	filterKeys := make(map[uint64]int)
	for _, categoryID := range req.CategoryIDs {
		filterKeys[categoryID] = 1
	}
	for _, modelCategory := range *modelCategories {
		categoryID := modelCategory.CategoryID
		if filterKeys[categoryID] == 0 {
			filterKeys[categoryID] = 2
		} else {
			filterKeys[categoryID] = 3
		}
	}
	for categoryID, key := range filterKeys {
		switch key {
		case 1:
			service.Create(categoryID, productID)
		case 2:
			service.Delete(categoryID)
		}
	}
	cateRepo := repositories.NewRepositoryCategory(service.DB)
	cateRepo.ReadByProductID(modelCategories, productID)
}
