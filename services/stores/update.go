package storesvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateBackOrder(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.IsBackOrder == 0 {
		modelStore.IsBackOrder = 1
	} else {
		modelStore.IsBackOrder = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateShowOutOfStockStatus(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.ShowOutOfStockStatus == 0 {
		modelStore.ShowOutOfStockStatus = 1
	} else {
		modelStore.ShowOutOfStockStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateShowStockLevelStatus(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.ShowStockLevelStatus == 0 {
		modelStore.ShowStockLevelStatus = 1
	} else {
		modelStore.ShowStockLevelStatus = 0
	}
	return service.DB.Save(modelStore).Error
}
