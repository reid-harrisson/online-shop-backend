package contsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelContents *[]models.ProductContentsWithTitle, req *requests.RequestProductContent, productID uint64) {
	filterKeys := make(map[uint64]int)
	for _, modelContent := range *modelContents {
		filterKeys[modelContent.ContentID] = 1
	}
	for _, contentID := range req.ContentIDs {
		if filterKeys[contentID] == 1 {
			filterKeys[contentID] = 3
		} else {
			filterKeys[contentID] = 2
		}
	}

	for contentID, key := range filterKeys {
		if key == 1 {
			service.Delete(contentID)
		} else if key == 2 {
			service.Create(contentID, productID)
		}
	}

	contRepo := repositories.NewRepositoryProductContent(service.DB)
	contRepo.ReadByProductID(modelContents, productID)
}
