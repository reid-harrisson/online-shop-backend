package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelTags *[]models.ProductTagsWithName, req *requests.RequestTag, modelProduct *models.Products) {
	filterKeys := make(map[string]int)
	for _, modelTag := range *modelTags {
		filterKeys[modelTag.TagName] = 1
	}
	for _, tag := range req.Tags {
		if filterKeys[tag] == 1 {
			filterKeys[tag] = 3
		} else {
			filterKeys[tag] = 2
		}
	}

	for tag, key := range filterKeys {
		if key == 1 {
			service.Delete(tag)
		} else if key == 2 {
			service.Create(tag, modelProduct)
		}
	}

	tagRepo := repositories.NewRepositoryTag(service.DB)
	tagRepo.ReadByProductID(modelTags, uint64(modelProduct.ID))
}
