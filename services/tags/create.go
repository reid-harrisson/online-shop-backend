package tagsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(tag string, modelTag *models.StoreTags) {
	modelTag.Name = tag
	service.DB.Where("name = ?", tag).First(modelTag)
	service.DB.Save(modelTag)
}
