package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.OwnerID = req.OwnerID
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.IsBackOrder = req.IsBackOrder
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms
	modelStore.Active = req.Active

	return service.DB.Create(&modelStore).Error
}
