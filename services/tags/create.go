package tagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelTag *models.StoreTags, tag string, storeID uint64) {
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
