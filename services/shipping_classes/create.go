package classsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingClass, modelClass *models.ShippingClasses) error {
	modelClass.StoreID = storeID
	modelClass.Name = req.Name
	modelClass.Description = req.Description
	modelClass.Priority = req.Priority

	if err := service.DB.Where("name = ?", req.Name).FirstOrCreate(&modelClass).Error; err != nil {
		return err
	}

	if modelClass.Description != req.Description {
		return service.DB.Model(models.ShippingClasses{}).Update("description", req.Description).Error
	}
	if modelClass.Priority != req.Priority {
		return service.DB.Model(models.ShippingClasses{}).Update("priority", req.Priority).Error
	}

	return nil
}
