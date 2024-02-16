package catesvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Update(modelCategory *models.BaseCategories, name string, parentID uint64) {
	if len(name) > 0 {
		modelCategory.Name = name
	}
	if parentID != 0 {
		modelCategory.ParentID = parentID
	}
	service.DB.Save(modelCategory)
}
