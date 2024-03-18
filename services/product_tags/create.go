package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	tagsvc "OnlineStoreBackend/services/tags"
	"fmt"
)

func (service *Service) Create(tag string, modelProduct *models.Products) {
	modelTag := models.StoreTags{}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	tagRepo.ReadByName(&modelTag, tag, modelProduct.StoreID)
	if modelTag.ID == 0 {
		tagService := tagsvc.NewServiceTag(service.DB)
		tagService.Create(tag, &modelTag, modelProduct.StoreID)
	}
	service.DB.Create(&models.ProductTags{
		TagID:     uint64(modelTag.ID),
		ProductID: uint64(modelProduct.ID),
	})
}

func (service *Service) CreateWithCSV(modelNewTags *[]models.ProductTags, tagMatches []string, tagIndices map[string]int) {
	modelCurTags := []models.ProductTags{}
	service.DB.Where("Concat(product_id, ':', tag_id) In (?)", tagMatches).Find(&modelCurTags)
	for _, modelTag := range modelCurTags {
		match := fmt.Sprintf("%d:%d", modelTag.ProductID, modelTag.TagID)
		index := tagIndices[match]
		(*modelNewTags)[index].ID = modelTag.ID
	}
	service.DB.Save(modelNewTags)
}
