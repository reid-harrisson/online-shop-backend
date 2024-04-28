package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelTags *[]models.ProductTagsWithName, req *requests.RequestProductTag, modelProduct *models.Products) error {
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
			err := service.Delete(tag)
			if err != nil {
				return err
			}
		} else if key == 2 {
			err := service.Create(tag, modelProduct)
			if err != nil {
				return err
			}
		}
	}

	tagRepo := repositories.NewRepositoryTag(service.DB)
	return tagRepo.ReadByProductID(modelTags, uint64(modelProduct.ID))
}
