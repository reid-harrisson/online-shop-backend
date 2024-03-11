package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
)

func (service *Service) UpdateStatus(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, orderStatus string) {
	status := utils.OrderStatusFromString(orderStatus)
	service.DB.Where("order_id = ?", orderID).Find(&modelItems)

	flagCompleted := true
	flagPending := true
	for index := range *modelItems {
		modelItem := &(*modelItems)[index]
		if modelItem.StoreID == storeID && modelItem.Status != status {
			modelItem.Status = status
			if status == utils.StatusOrderProcessing {
				modelVar := models.ProductVariations{}
				service.DB.First(&modelVar, modelItem.VariationID)
				modelVar.StockLevel -= modelItem.Quantity
				service.DB.Save(&modelVar)
			} else if status == utils.StatusOrderPending {
				modelVar := models.ProductVariations{}
				service.DB.First(&modelVar, modelItem.VariationID)
				modelVar.StockLevel += modelItem.Quantity
				service.DB.Save(&modelVar)
			}
			service.DB.Save(modelItem)
		}
		if modelItem.Status != utils.StatusOrderCompleted {
			flagCompleted = false
		}
		if modelItem.Status != utils.StatusOrderPending {
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

func (service *Service) UpdateCoupon(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, modelCoupon *models.Coupons) {

}
