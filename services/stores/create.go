package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.UserID = req.UserID
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockQuantity = req.ShowStockQuantity
	modelStore.ShowOutOfStockProducts = req.ShowOutOfStockProducts
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms
	modelStore.FlatRateShipping = req.FlatRateShipping
	modelStore.BackOrder = req.BackOrder
	modelStore.Active = req.Active
	service.DB.Create(&modelStore)
	return nil
}
