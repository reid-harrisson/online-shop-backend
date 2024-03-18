package tagsvc

import (
	"OnlineStoreBackend/models"
	"strings"
)

func (service *Service) Create(tag string, modelTag *models.StoreTags, storeID uint64) {
	modelTag.Name = tag
	modelTag.StoreID = storeID
	service.DB.Where("name = ?", tag).First(modelTag)
	service.DB.Save(modelTag)
}

func (service *Service) CreateWithCSV(modelNewTags *[]models.StoreTags, tagNames []string, tagIndices map[string]int) {
	modelCurTags := []models.StoreTags{}
	service.DB.Where("name In (?)", tagNames).Find(&modelCurTags)
	for _, modelTag := range modelCurTags {
		index := tagIndices[modelTag.Name] - 1
		(*modelNewTags)[index].ID = modelTag.ID
	}
	service.DB.Save(modelNewTags)
}

func (service *Service) CreateWithCSV1(modelTags *[]models.StoreTags, tags []string, storeID uint64) {
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
		if len(tags[i]) == 0 {
			tags = append(tags[:i], tags[i+1:]...)
		}
	}
	if len(tags) == 0 {
		return
	}
	service.DB.Where("name In (?)", tags).Find(modelTags)
	mapTag := make(map[string]int)
	for index, modelTag := range *modelTags {
		mapTag[modelTag.Name] = index + 1
	}
	for _, tag := range tags {
		if mapTag[tag] == 0 {
			modelTag := models.StoreTags{}
			modelTag.Name = tag
			modelTag.StoreID = storeID
			service.DB.Create(&modelTag)
			*modelTags = append(*modelTags, modelTag)
		}
	}
}
