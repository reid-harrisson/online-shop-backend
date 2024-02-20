package orditmsvc

import (
	"OnlineStoreBackend/models"
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
	service.DB.Where("id = ? And store_id = ?", orderID, storeID).
		Update("status", status)
}
