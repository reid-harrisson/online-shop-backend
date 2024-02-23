package shipmthsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingOption) error {
	modelMethod := models.ShippingMethods{
		StoreID: storeID,
		Method:  utils.ShippingMethodsFromString(req.Method),
	}
	return service.DB.Create(&modelMethod).Error
}
