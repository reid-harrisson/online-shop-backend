package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelCategory *models.StoreCategories, req *requests.RequestCategory, storeID uint64) {
	modelCategory.Name = req.Name
	modelCategory.StoreID = storeID
	modelCategory.ParentID = req.ParentID
	service.DB.Create(modelCategory)
}
