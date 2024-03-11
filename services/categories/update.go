package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelCategory *models.StoreCategories, req *requests.RequestCategory) {
	modelCategory.Name = req.Name
	modelCategory.ParentID = req.ParentID
	service.DB.Save(modelCategory)
}
