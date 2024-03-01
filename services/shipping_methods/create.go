package shipmthsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingMethod) error {
	modelMethod := models.ShippingMethods{
		StoreID:       storeID,
		Method:        utils.ShippingMethodsFromString(req.Method),
		FlatRate:      req.FlatRate,
		BaseRate:      req.BaseRate,
		RatePerItem:   req.RatePerItem,
		RatePerWeight: req.RatePerWeight,
		RatePerTotal:  req.RatePerTotal,
	}
	return service.DB.Create(&modelMethod).Error
}
