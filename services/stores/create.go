package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStore *models.Stores, req *requests.RequestStore, userID uint64) error {
	modelAddr := models.Users{}
	addrRepo := repositories.NewRepositoryUser(service.DB)
	err := addrRepo.ReadByID(&modelAddr, userID)
	if err != nil {
		return err
	}

	modelStore.CompanyID = modelAddr.CompanyID

	modelStore.OwnerID = userID
	modelStore.Name = req.Name
	modelStore.ContactPhone = req.ContactPhone
	modelStore.ContactEmail = req.ContactEmail
	modelStore.ShowStockLevelStatus = req.ShowStockLevelStatus
	modelStore.ShowOutOfStockStatus = req.ShowOutOfStockStatus
	modelStore.DeliveryPolicy = req.DeliveryPolicy
	modelStore.ReturnsPolicy = req.ReturnsPolicy
	modelStore.Terms = req.Terms

	return service.DB.Create(&modelStore).Error
}
