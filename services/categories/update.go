package catesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(id uint64, modelCategory *models.Categories, req *requests.RequestCategory) error {
	err := service.DB.Where("id = ?", id).First(&modelCategory).Error
	if err != nil {
		return nil
	}

	modelCategory.Name = req.Name
	modelCategory.ParentID = req.ParentID
	return service.DB.Save(modelCategory).Error
}
