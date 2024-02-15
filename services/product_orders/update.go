package ordsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateStatus(modelOrders *[]models.ProductOrders, req *requests.RequestProductOrderStatus, id uint64) error {
	service.DB.Where("id = ?", id).Find(modelOrders)
	for _, modelOrder := range *modelOrders {
		service.DB.Model(&modelOrder).Update("Status", req.Status)
		modelOrder.Status = req.Status
	}
	return nil
}

func (service *Service) UpdateFreeShipping(modelOrders *[]models.ProductOrders, storeID uint64, miniumAmount float64, productID uint64) error {
	service.DB.Where("? Is Null")
	return nil
}
