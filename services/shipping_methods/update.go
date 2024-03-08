package methsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	flatsvc "OnlineStoreBackend/services/shipping_flat_rates"
)

func (service *Service) UpdateShippingLocalPickup(req *requests.RequestShippingLocalPickup, modelMethod *models.ShippingMethods) error {
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.Cost = req.Cost
	return service.DB.Save(modelMethod).Error
}

func (service *Service) UpdateShippingFree(req *requests.RequestShippingFree, modelMethod *models.ShippingMethods) error {
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	modelMethod.ZoneID = req.ZoneID
	modelMethod.Requirement = req.Requirement
	modelMethod.MinimumOrderAmount = req.MinimumOrderAmount
	return service.DB.Save(modelMethod).Error
}

func (service *Service) UpdateShippingFlatRate(req *requests.RequestShippingFlatRate, modelMethod *models.ShippingMethods, modelRates *[]models.ShippingFlatRates) error {
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.Cost = req.Cost
	if err := service.DB.Save(modelMethod).Error; err != nil {
		return err
	}
	flatService := flatsvc.NewServiceShippingFlatRate(service.DB)
	return flatService.Create(uint64(modelMethod.ID), req.Rates, modelRates)
}
