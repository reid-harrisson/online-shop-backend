package orditmsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) UpdateStatus(storeID uint64, orderID uint64, orderStatus string) {
	status := utils.OrderStatusFromString(orderStatus)
	service.DB.Model(models.OrderItems{}).Where("id = ? And store_id = ?", orderID, storeID).
		Update("status", status)
}

func (service *Service) UpdateShippingMethod(storeID uint64, orderID uint64, method string) error {
	return service.DB.Model(models.OrderItems{}).Where("order_id = ? And store_id = ?", orderID, storeID).
		Update("shipping_method", utils.ShippingMethodsFromString(method)).Error
}
