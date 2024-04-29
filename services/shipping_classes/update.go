package classsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(req *requests.RequestShippingClass, modelClass *models.ShippingClasses) error {
	modelClass.Name = req.Name
	modelClass.Description = req.Description
	modelClass.Priority = req.Priority

	return service.DB.Save(modelClass).Error
}
