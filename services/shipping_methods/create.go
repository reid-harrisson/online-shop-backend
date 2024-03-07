package shipmthsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingMethod) error {
	modelMethod := models.ShippingMethods{
		StoreID:             storeID,
		Method:              utils.ShippingMethodsFromString(req.Method),
		Requirement:         req.Requirement,
		MinimumOrderAmount:  req.MinimumOrderAmount,
		TaxStatus:           req.TaxStatus,
		Cost:                req.Cost,
		TaxIncluded:         req.TaxIncluded,
		HandlingFee:         req.HandlingFee,
		MaximumShippingCost: req.MaximumShippingCost,
		CalculationType:     req.CalculationType,
		HandlingFeePerClass: req.HandlingFeePerClass,
		MinimumCostPerClass: req.MinimumCostPerClass,
		MaximumCostPerClass: req.MaximumCostPerClass,
		DiscountInMinMax:    req.DiscountInMinMax,
		TaxInMinMax:         req.TaxInMinMax,
	}
	return service.DB.Create(&modelMethod).Error
}
