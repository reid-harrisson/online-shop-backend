package prodcatesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelCategories *[]models.ProductCategoriesWithName, req *requests.RequestProductCategory, productID uint64) error {
	filterKeys := make(map[uint64]int)
	for _, modelCategory := range *modelCategories {
		filterKeys[modelCategory.CategoryID] = 1
	}
	for _, categoryID := range req.CategoryIDs {
		if filterKeys[categoryID] == 1 {
			filterKeys[categoryID] = 3
		} else {
			filterKeys[categoryID] = 2
		}
	}

	for categoryID, key := range filterKeys {
		if key == 1 {
			service.Delete(categoryID)
		} else if key == 2 {
			service.Create(categoryID, productID)
		}
	}

	cateRepo := repositories.NewRepositoryCategory(service.DB)

	return cateRepo.ReadByProductID(modelCategories, productID)
}
