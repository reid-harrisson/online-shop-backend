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
	modelStore.Active = req.Active
	modelStore.BackgroundColor1 = req.BackgroundColor1
	modelStore.BackgroundColor2 = req.BackgroundColor2
	modelStore.StoreBackground = req.StoreBackground
	modelStore.StoreLogo = req.StoreLogo
	modelStore.Description = req.Description
	modelStore.HeaderLayoutStyle = req.HeaderLayoutStyle
	modelStore.ShowStoreLogo = req.ShowStoreLogo
	modelStore.ShowStoreTitleText = req.ShowStoreTitleText
	modelStore.Website = req.Website
	modelStore.WebsiteButtonColor = req.WebsiteButtonColor

	return service.DB.Create(&modelStore).Error
}
