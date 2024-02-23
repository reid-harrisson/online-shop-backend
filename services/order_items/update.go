package orditmsvc

import (
	"OnlineStoreBackend/models"
	"strings"
)

func (service *Service) UpdateStatus(storeID uint64, orderID uint64, orderStatus string) {
	status := models.StatusOrderPending
	switch orderStatus {
	case "Pending":
		status = models.StatusOrderPending
	case "Active":
		status = models.StatusOrderProcessing
	case "Cancelled":
		status = models.StatusOrderPaid
	case "Expired":
		status = models.StatusOrderShipping
	case "Overdue":
		status = models.StatusOrderCompleted
	}
	service.DB.Model(models.OrderItems{}).Where("id = ? And store_id = ?", orderID, storeID).
		Update("status", status)
}

func (service *Service) UpdateShippingMethod(storeID uint64, orderID uint64, shippingMethod string) error {
	method := models.MethodPickUp
	switch strings.ToLower(shippingMethod) {
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
	return service.DB.Model(models.OrderItems{}).Where("order_id = ? And store_id = ?", orderID, storeID).
		Update("shipping_method", method).Error
}
