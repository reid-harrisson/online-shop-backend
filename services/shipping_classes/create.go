package classsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingClass, modelClass *models.ShippingClasses) {
	modelClass.StoreID = storeID
	modelClass.Name = req.Name
	modelClass.Description = req.Description
	modelClass.Priority = req.Priority

	service.DB.Where("name = ?", req.Name).FirstOrCreate(&modelClass)
	if modelClass.Description != req.Description {
		service.DB.Model(models.ShippingClasses{}).Update("description", req.Description)
	}
	if modelClass.Priority != req.Priority {
		service.DB.Model(models.ShippingClasses{}).Update("priority", req.Priority)
	}
}
