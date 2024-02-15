package storesvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateBackOrder(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)
	if modelStore.BackOrder == 0 {
		modelStore.BackOrder = 1
	} else {
		modelStore.BackOrder = 0
	}
	service.DB.Save(modelStore)
	return nil
}

func (service *Service) UpdateStockTracking(storeID uint64, modelStore *models.Stores) error {
	service.DB.First(modelStore, storeID)
	if modelStore.ShowOutOfStockProducts == 0 {
		modelStore.ShowOutOfStockProducts = 1
	} else {
		modelStore.ShowOutOfStockProducts = 0
	}
	service.DB.Save(modelStore)
	return nil
}
