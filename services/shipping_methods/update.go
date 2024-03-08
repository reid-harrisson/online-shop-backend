package methsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	flatsvc "OnlineStoreBackend/services/shipping_flat_rates"
	tablesvc "OnlineStoreBackend/services/shipping_table_rates"
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

func (service *Service) UpdateShippingTableRate(req *requests.RequestShippingTableRate, modelMethod *models.ShippingMethods, modelRates *[]models.ShippingTableRates) error {
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.TaxIncluded = req.TaxIncluded
	modelMethod.HandlingFee = req.HandlingFee
	modelMethod.MaximumShippingCost = req.MaximumShippingCost
	modelMethod.CalculationType = req.CalculationType
	modelMethod.HandlingFeePerClass = req.HandlingFeePerClass
	modelMethod.MinimumCostPerClass = req.MinimumCostPerClass
	modelMethod.MaximumCostPerClass = req.MaximumCostPerClass
	modelMethod.DiscountInMinMax = req.DiscountInMinMax
	modelMethod.TaxInMinMax = req.TaxInMinMax
	if err := service.DB.Save(modelMethod).Error; err != nil {
		return err
	}
	tableService := tablesvc.NewServiceShippingTableRate(service.DB)
	return tableService.Create(uint64(modelMethod.ID), req.Rates, modelRates)
}
