package classsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(req *requests.RequestShippingClass, modelClass *models.ShippingClasses) {
	modelClass.Name = req.Name
	modelClass.Description = req.Description
	modelClass.Priority = req.Priority

	service.DB.Save(modelClass)
}
