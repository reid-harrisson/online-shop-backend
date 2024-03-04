package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	tagsvc "OnlineStoreBackend/services/tags"
)

func (service *Service) Create(tag string, modelProduct *models.Products) {
	modelTag := models.StoreTags{}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	tagRepo.ReadByName(&modelTag, tag)
	if modelTag.ID == 0 {
		tagService := tagsvc.NewServiceTag(service.DB)
		tagService.Create(tag, &modelTag, modelProduct.StoreID)
	}
	service.DB.Create(&models.ProductTags{
		TagID:     uint64(modelTag.ID),
		ProductID: uint64(modelProduct.ID),
	})
}

func (service *Service) CreateWithCSV(modelTags []models.StoreTags, productID uint64) {
	for _, modelTag := range modelTags {
		service.DB.Create(&models.ProductTags{
			TagID:     uint64(modelTag.ID),
			ProductID: productID,
		})
	}
}
