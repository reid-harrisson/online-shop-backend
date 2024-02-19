package storesvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateBackOrder(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)

	if modelStore.IsBackOrder == 0 {
		modelStore.IsBackOrder = 1
	} else {
		modelStore.IsBackOrder = 0
	}
	service.DB.Save(modelStore)
	return nil
}

func (service *Service) UpdateStockTracking(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)
	if modelStore.ShowOutOfStockStatus == 0 {
		modelStore.ShowOutOfStockStatus = 1
	} else {
		modelStore.ShowOutOfStockStatus = 0
	}
	service.DB.Save(modelStore)
	return nil
}
