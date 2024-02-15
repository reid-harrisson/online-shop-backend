package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelTags *[]models.ProductTagsWithName, req *requests.RequestTag, productID uint64) {
	filterKeys := make(map[uint64]int)
	for _, tagID := range req.TagIDs {
		filterKeys[tagID] = 1
	}
	for _, modelTag := range *modelTags {
		tagID := modelTag.TagID
		if filterKeys[tagID] == 0 {
			filterKeys[tagID] = 2
		} else {
			filterKeys[tagID] = 3
		}
	}
	for tagID, key := range filterKeys {
		switch key {
		case 1:
			service.Create(tagID, productID)
		case 2:
			service.Delete(tagID)
		}
	}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	tagRepo.ReadByProductID(modelTags, productID)
}
