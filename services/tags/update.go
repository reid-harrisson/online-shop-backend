package tagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelTag *models.Tags, req *requests.RequestTag) {
	modelTag.Name = req.Name
	service.DB.Save(modelTag)
}
