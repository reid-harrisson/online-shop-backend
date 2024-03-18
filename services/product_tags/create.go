package prodtagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	tagsvc "OnlineStoreBackend/services/tags"
	"fmt"
	"sort"
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

func (service *Service) CreateWithCSV(prodTags map[uint64][]uint64) {
	modelNewTags := []models.ProductTags{}
	modelCurTags := []models.ProductTags{}
	matches := []string{}
	indices := map[string]int{}
	for prodID, tagIDs := range prodTags {
		for _, tagID := range tagIDs {
			modelNewTags = append(modelNewTags, models.ProductTags{
				ProductID: prodID,
				TagID:     tagID,
			})
			match := fmt.Sprintf("%d:%d", prodID, tagID)
			if indices[match] == 0 {
				matches = append(matches, match)
				indices[match] = len(matches)
			}
		}
	}
	sort.Slice(modelNewTags, func(i, j int) bool {
		if modelNewTags[i].ProductID == modelNewTags[j].ProductID {
			return modelNewTags[i].TagID < modelNewTags[j].TagID
		}
		return modelNewTags[i].ProductID < modelNewTags[j].ProductID
	})
	service.DB.Where("Concat(product_id, ':', tag_id) In (?)", matches).Find(&modelCurTags)
	for _, modelTag := range modelCurTags {
		match := fmt.Sprintf("%d:%d", modelTag.ProductID, modelTag.TagID)
		index := indices[match] - 1
		modelNewTags[index].ID = modelTag.ID
	}
	service.DB.Save(&modelNewTags)
}

func (service *Service) CreateWithCSV1(modelTags []models.StoreTags, productID uint64) {
	for _, modelTag := range modelTags {
		service.DB.Create(&models.ProductTags{
			TagID:     uint64(modelTag.ID),
			ProductID: productID,
		})
	}
}
