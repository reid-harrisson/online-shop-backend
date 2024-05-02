package tagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(id uint64, modelTag *models.Tags, req *requests.RequestTag) error {
	err := service.DB.Where("id = ?", id).First(modelTag).Error
	if err != nil {
		return err
	}

	modelTag.Name = req.Name
	return service.DB.Save(modelTag).Error
}
