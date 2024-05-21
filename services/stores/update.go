package storesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateOutOfStockStatus(storeID uint64, modelStore *models.Stores) error {
	if err := service.DB.First(modelStore, storeID).Error; err != nil {
		return err
	}

	if modelStore.ShowOutOfStockStatus == 0 {
		modelStore.ShowOutOfStockStatus = 1
	} else {
		modelStore.ShowOutOfStockStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateStockLevelStatus(storeID uint64, modelStore *models.Stores) error {
	if err := service.DB.First(modelStore, storeID).Error; err != nil {
		return err
	}

	if modelStore.ShowStockLevelStatus == 0 {
		modelStore.ShowStockLevelStatus = 1
	} else {
		modelStore.ShowStockLevelStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) Update(modelStore *models.Stores, req *requests.RequestStore, storeID uint64) error {
	if err := service.DB.First(modelStore, storeID).Error; err != nil {
		return err
	}

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

	return service.DB.Save(modelStore).Error
}
