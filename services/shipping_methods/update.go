package methsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateShippingLocalPickup(req *requests.RequestShippingLocalPickup, modelMethod *models.ShippingMethods) error {
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.Cost = req.Cost
	return service.DB.Save(modelMethod).Error
}

func (service *Service) UpdateShippingFree(req *requests.RequestShippingFree, modelMethod *models.ShippingMethods) error {
	modelMethod.ZoneID = req.ZoneID
	modelMethod.Requirement = req.Requirement
	modelMethod.MinimumOrderAmount = req.MinimumOrderAmount
	return service.DB.Save(modelMethod).Error
}
