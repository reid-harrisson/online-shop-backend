package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
)

func (service *Service) UpdateStatus(storeID uint64, orderID uint64, orderStatus string) {
	status := models.StatusOrderPending
	switch orderStatus {
	case "Pending":
		status = models.StatusOrderPending
	case "Payment Processing":
		status = models.StatusOrderPaymentProcessing
	case "Paid":
		status = models.StatusOrderPaid
	case "Processing":
		status = models.StatusOrderProcessing
	case "Shipping Processing":
		status = models.StatusOrderShippingProcessing
	case "Shipping":
		status = models.StatusOrderShipping
	case "Shipped":
		status = models.StatusOrderShipped
	case "Completed":
		status = models.StatusOrderCompleted
	}
	service.DB.Table("store_order_items").
		Where("order_id = ? And store_id = ?", orderID, storeID).
		Update("status", status)

	modelOrders := make([]models.CustomerOrdersWithDetail, 0)
	orderRepo := repositories.NewRepositoryOrder(service.DB)
	orderRepo.ReadByOrderID(&modelOrders, orderID)
	flagCompleted := true
	flagPending := true
	for _, modelOrder := range modelOrders {
		if modelOrder.ProductStatus != models.StatusOrderCompleted {
			flagCompleted = false
		}
		if modelOrder.ProductStatus != models.StatusOrderPending {
			flagPending = false
		}
	}
	if flagCompleted {
		status = models.StatusOrderCompleted
	} else if flagPending {
		status = models.StatusOrderPending
	} else {
		status = models.StatusOrderProcessing
	}
	service.DB.Table("store_orders").
		Where("id = ?", orderID).
		Update("status", status)
}

func (service *Service) UpdateBillingAddress(orderID uint64, addressID uint64) {
	service.DB.Model(models.Orders{}).Update("billing_address_id", addressID)
}

func (service *Service) UpdateShippingAddress(orderID uint64, addressID uint64) {
	service.DB.Model(models.Orders{}).Update("shipping_address_id", addressID)
}
