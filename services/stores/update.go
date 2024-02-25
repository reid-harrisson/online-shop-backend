package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateBackOrder(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.BackOrderStatus == 0 {
		modelStore.BackOrderStatus = 1
	} else {
		modelStore.BackOrderStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateOutOfStockStatus(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.ShowOutOfStockStatus == 0 {
		modelStore.ShowOutOfStockStatus = 1
	} else {
		modelStore.ShowOutOfStockStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateStockLevelStatus(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.ShowStockLevelStatus == 0 {
		modelStore.ShowStockLevelStatus = 1
	} else {
		modelStore.ShowStockLevelStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) Update(modelStore *models.Stores, req *requests.RequestStore) error {
	modelStore.CompanyID = req.CompanyID
	modelStore.OwnerID = req.OwnerID
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.BackOrderStatus = req.BackOrderStatus
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ContactPhone = req.ContactPhone

	return service.DB.Save(modelStore).Error
}
