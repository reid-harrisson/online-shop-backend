package methsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	flatsvc "OnlineStoreBackend/services/shipping_flat_rates"
	tablesvc "OnlineStoreBackend/services/shipping_table_rates"
)

func (service *Service) CreateShippingLocalPickup(storeID uint64, req *requests.RequestShippingLocalPickup, modelMethod *models.ShippingMethods) error {
	err := service.DB.Where("title = ?", req.Title).First(modelMethod).Error
	modelMethod.Method = utils.PickUp
	modelMethod.StoreID = storeID
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.Cost = req.Cost
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	if err != nil {
		return service.DB.Create(modelMethod).Error
	}
	return service.DB.Save(modelMethod).Error
}

func (service *Service) CreateShippingFree(storeID uint64, req *requests.RequestShippingFree, modelMethod *models.ShippingMethods) error {
	err := service.DB.Where("title = ?", req.Title).First(modelMethod).Error
	modelMethod.Method = utils.FreeShipping
	modelMethod.StoreID = storeID
	modelMethod.ZoneID = req.ZoneID
	modelMethod.Requirement = req.Requirement
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	modelMethod.MinimumOrderAmount = req.MinimumOrderAmount
	if err != nil {
		return service.DB.Create(modelMethod).Error
	}
	return service.DB.Save(modelMethod).Error
}

func (service *Service) CreateShippingFlatRate(storeID uint64, req *requests.RequestShippingFlatRate, modelMethod *models.ShippingMethods, modelRates *[]models.ShippingFlatRates) error {
	err := service.DB.Where("title = ?", req.Title).First(modelMethod).Error
	modelMethod.Method = utils.FlatRate
	modelMethod.StoreID = storeID
	modelMethod.ZoneID = req.ZoneID
	modelMethod.TaxStatus = req.TaxStatus
	modelMethod.Cost = req.Cost
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	if err != nil {
		return service.DB.Create(modelMethod).Error
	}
	if err := service.DB.Save(modelMethod).Error; err != nil {
		return err
	}
	flatService := flatsvc.NewServiceShippingFlatRate(service.DB)
	return flatService.Create(uint64(modelMethod.ID), req.Rates, modelRates)
}

func (service *Service) CreateShippingTableRate(storeID uint64, req *requests.RequestShippingTableRate, modelMethod *models.ShippingMethods, modelRates *[]models.ShippingTableRates) error {
	err := service.DB.Where("title = ?", req.Title).First(modelMethod).Error
	modelMethod.Method = utils.TableRate
	modelMethod.StoreID = storeID
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
	modelMethod.Title = req.Title
	modelMethod.Description = req.Description
	if err != nil {
		return service.DB.Create(modelMethod).Error
	}
	if err := service.DB.Save(modelMethod).Error; err != nil {
		return err
	}
	tableService := tablesvc.NewServiceShippingTableRate(service.DB)
	return tableService.CreateMany(uint64(modelMethod.ID), req.Rates, modelRates)
}
