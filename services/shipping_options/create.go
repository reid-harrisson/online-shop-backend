package shipoptsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strings"
)

func (service *Service) Create(storeID uint64, req *requests.RequestShippingOption) error {
	method := models.MethodPickUp
	switch strings.ToLower(req.Method) {
	case "pick up":
		method = models.MethodPickUp
	case "flat rate":
		method = models.MethodFlatRate
	case "table rate":
		method = models.MethodTableRate
	case "free shipping":
		method = models.MethodFreeShipping
	case "real time":
		method = models.MethodRealTime
	}
	modelMethod := models.ShippingOptions{
		StoreID: storeID,
		Method:  method,
	}
	return service.DB.Create(&modelMethod).Error
}
