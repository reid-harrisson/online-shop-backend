package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
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

func (service *Service) UpdateOrderItemStatus(orderID uint64, status string) {
	service.DB.
		Model(models.OrderItems{}).
		Where("order_id = ?", orderID).
		Update("status", utils.OrderStatusFromString(status))
}

func (service *Service) UpdateBillingAddress(orderID uint64, addressID uint64) {
	service.DB.
		Model(models.Orders{}).
		Update("billing_address_id", addressID)
}

func (service *Service) UpdateShippingAddress(orderID uint64, addressID uint64) {
	service.DB.
		Model(models.Orders{}).
		Update("shipping_address_id", addressID)
}

func (service *Service) UpdateCoupon(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, modelCoupon *models.Coupons) error {
	service.DB.Where("order_id = ? And store_id = ?", orderID, storeID).Find(modelItems)
	size := len(*modelItems)
	if size == 0 {
		return nil
	}

	modelRates := []models.ShippingTableRates{}
	methRepo := repositories.NewRepositoryShippingMethod(service.DB)
	methRepo.ReadRates(&modelRates, storeID)

	for index := range *modelItems {
		modelShip := models.ShippingData{}
		shipRepo := repositories.NewRepositoryShippingData(service.DB)
		shipRepo.ReadByVariationID(&modelShip, (*modelItems)[index].VariationID)
		if modelCoupon.DiscountType == utils.FixedProductDiscount {
			(*modelItems)[index].Price -= modelCoupon.CouponAmount
			(*modelItems)[index].SubTotalPrice = (*modelItems)[index].Price * (*modelItems)[index].Quantity
		} else if modelCoupon.DiscountType == utils.FixedCartDiscount {
			(*modelItems)[index].SubTotalPrice -= modelCoupon.CouponAmount / float64(size)
		} else if modelCoupon.DiscountType == utils.PercentageDiscount {
			(*modelItems)[index].SubTotalPrice *= (100 - modelCoupon.CouponAmount) / 100
		}
		(*modelItems)[index].TaxAmount = (*modelItems)[index].SubTotalPrice * (*modelItems)[index].TaxRate / 100
		(*modelItems)[index].ShippingPrice = GetShippingPrice(modelRates, (*modelItems)[index].SubTotalPrice, (*modelItems)[index].Quantity, modelShip)
		(*modelItems)[index].TotalPrice = (*modelItems)[index].ShippingPrice + (*modelItems)[index].SubTotalPrice + (*modelItems)[index].TaxAmount
	}

	service.DB.Delete(modelCoupon)
	return service.DB.Save(modelItems).Error
}
