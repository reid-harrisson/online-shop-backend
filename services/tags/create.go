package tagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelTag *models.Tags, tag string, storeID uint64) error {
	err := service.DB.Where("name = ? And store_id = ?", tag, storeID).First(modelTag).Error
	if err != nil {
		return err
	}

	modelTag.Name = tag
	modelTag.StoreID = storeID

	return service.DB.Create(modelTag).Error
}

func (service *Service) CreateWithCSV(modelNewTags *[]models.Tags, tagNames []string, tagIndices map[string]int) error {
	modelCurTags := []models.Tags{}
	if err := service.DB.Where("name In (?)", tagNames).Find(&modelCurTags).Error; err != nil {
		return nil
	}
	for _, modelTag := range modelCurTags {
		index := tagIndices[modelTag.Name] - 1
		(*modelNewTags)[index].ID = modelTag.ID
	}
	return service.DB.Save(modelNewTags).Error
}
