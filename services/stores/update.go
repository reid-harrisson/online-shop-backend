package storesvc

import "OnlineStoreBackend/models"

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

	if modelStore.OutOfStockStatus == 0 {
		modelStore.OutOfStockStatus = 1
	} else {
		modelStore.OutOfStockStatus = 0
	}
	return service.DB.Save(modelStore).Error
}

func (service *Service) UpdateStockLevelStatus(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.StockLevelStatus == 0 {
		modelStore.StockLevelStatus = 1
	} else {
		modelStore.StockLevelStatus = 0
	}
	return service.DB.Save(modelStore).Error
}
