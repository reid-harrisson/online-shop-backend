package tagsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(modelTag *models.StoreTags, req *requests.RequestTag, storeID uint64) {
	modelTag.Name = req.Name
	modelTag.StoreID = storeID
	service.DB.Save(modelTag)
}
