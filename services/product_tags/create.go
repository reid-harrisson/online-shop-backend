package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	tagsvc "OnlineStoreBackend/services/tags"
	"fmt"
)

func (service *Service) Create(tag string, modelProduct *models.Products) error {
	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(service.DB)
	err := tagRepo.ReadByNameAndStoreID(&modelTag, tag, modelProduct.StoreID)
	if err != nil {
		return err
	}

	if modelTag.ID == 0 {
		tagService := tagsvc.NewServiceTag(service.DB)
		err = tagService.Create(&modelTag, tag, modelProduct.StoreID)
		if err != nil {
			return err
		}
	}

	return service.DB.Create(&models.ProductTags{
		TagID:     uint64(modelTag.ID),
		ProductID: uint64(modelProduct.ID),
	}).Error
}

func (service *Service) CreateWithCSV(modelNewTags *[]models.ProductTags, tagMatches []string, tagIndices map[string]int) error {
	modelCurTags := []models.ProductTags{}
	if err := service.DB.Where("Concat(product_id, ':', tag_id) In (?)", tagMatches).Find(&modelCurTags).Error; err != nil {
		return err
	}
	for _, modelTag := range modelCurTags {
		match := fmt.Sprintf("%d:%d", modelTag.ProductID, modelTag.TagID)
		index := tagIndices[match]
		(*modelNewTags)[index].ID = modelTag.ID
	}
	return service.DB.Save(modelNewTags).Error
}
