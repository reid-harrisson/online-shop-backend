package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore, userID uint64) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.OwnerID = userID
	modelStore.Name = req.Name
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ContactPhone = req.ContactPhone

	return service.DB.Create(&modelStore).Error
}
