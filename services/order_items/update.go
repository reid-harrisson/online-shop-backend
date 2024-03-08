package orditmsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
)

func (service *Service) UpdateStatus(storeID uint64, orderID uint64, orderStatus string) {
	status := utils.OrderStatusFromString(orderStatus)
	service.DB.Model(models.OrderItems{}).Where("id = ? And store_id = ?", orderID, storeID).
		Update("status", status)
}

func (service *Service) UpdateShippingMethod(modelItems *[]models.OrderItems, storeID uint64, orderID uint64, methodID uint64) {
	service.DB.Where("order_id = ? And store_id = ?", orderID, storeID).Find(modelItems)
	shipRepo := repositories.NewRepositoryShippingData(service.DB)
	methRepo := repositories.NewRepositoryShippingMethod(service.DB)
	for index := range *modelItems {
		modelMethod := models.ShippingMethods{}
		methRepo.ReadByID(&modelMethod, methodID)
		modelShip := models.ShippingData{}
		shipRepo.ReadByVariationID(&modelShip, (*modelItems)[index].VariationID)
		totalPrice := (*modelItems)[index].SubTotalPrice
		shippingPrice := float64(0)
		if modelMethod.ID != 0 {
			switch modelMethod.Method {
			case utils.FlatRate:
				shippingPrice = 0
			case utils.TableRate:
				shippingPrice = 0
			}
		}
		(*modelItems)[index].ShippingMethodID = methodID
		(*modelItems)[index].ShippingPrice = shippingPrice
		(*modelItems)[index].TotalPrice = totalPrice + (*modelItems)[index].TaxAmount + shippingPrice
		service.DB.Save(&(*modelItems)[index])
	}
}
