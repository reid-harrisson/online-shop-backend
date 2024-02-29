package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strings"
)

func (service *Service) Create(modelCategory *models.StoreCategories, req *requests.RequestCategory, storeID uint64) {
	modelCategory.Name = req.Name
	modelCategory.StoreID = storeID
	modelCategory.ParentID = req.ParentID
	service.DB.Create(modelCategory)
}

func (service *Service) CreateWithCSV(modelCategories *[]models.StoreCategories, categories []string, storeID uint64) {
	for i := range categories {
		categories[i] = strings.TrimSpace(categories[i])
		if len(categories[i]) == 0 {
			categories = append(categories[:i], categories[i+1:]...)
		}
	}
	if len(categories) == 0 {
		return
	}
	service.DB.Where("name In (?)", categories).Find(modelCategories)
	mapCategory := make(map[string]int)
	for index, modelCategory := range *modelCategories {
		mapCategory[modelCategory.Name] = index + 1
	}
	preIndex := 0
	for _, category := range categories {
		index := mapCategory[category]
		if index > 0 {
			if preIndex > 0 {
				if (*modelCategories)[index-1].ParentID != uint64((*modelCategories)[preIndex-1].ID) {
					service.DB.Model(models.StoreCategories{}).Where("id = ?", index-1).Update("parent_id", uint64((*modelCategories)[preIndex-1].ID))
				}
			} else {
				if (*modelCategories)[index-1].ParentID != 0 {
					service.DB.Model(models.StoreCategories{}).Where("id = ?", index-1).Update("parent_id", 0)
				}
			}
		} else {
			modelCategory := models.StoreCategories{}
			modelCategory.Name = category
			modelCategory.StoreID = storeID
			if preIndex > 0 {
				modelCategory.ParentID = uint64((*modelCategories)[preIndex-1].ID)
			} else {
				modelCategory.ParentID = 0
			}
			service.DB.Create(&modelCategory)
			mapCategory[category] = len(*modelCategories)
			*modelCategories = append(*modelCategories, modelCategory)
		}
		preIndex = index
	}
}
