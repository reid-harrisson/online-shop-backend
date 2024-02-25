package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
)

func (service *Service) UpdateStatus(storeID uint64, orderID uint64, orderStatus string) {
	status := utils.OrderStatusFromString(orderStatus)
	service.DB.Table("store_order_items").
		Where("order_id = ? And store_id = ?", orderID, storeID).
		Update("status", status)

	modelOrder := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(service.DB)
	orderRepo.ReadByOrderID(&modelOrder, orderID)
	flagCompleted := true
	flagPending := true
	for _, modelItem := range modelOrder.Items {
		if modelItem.ProductStatus != utils.StatusOrderCompleted {
			flagCompleted = false
		}
		if modelItem.ProductStatus != utils.StatusOrderPending {
			flagPending = false
		}
	}
	if flagCompleted {
		status = utils.StatusOrderCompleted
	} else if flagPending {
		status = utils.StatusOrderPending
	} else {
		status = utils.StatusOrderProcessing
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
