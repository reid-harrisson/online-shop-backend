package prodvarsvc

import "OnlineStoreBackend/models"

func (service *Service) UpdateStockLevel(modelVar *models.ProductVariations, stockLevel float64) {
	modelVar.StockLevel = stockLevel
	service.DB.Save(modelVar)
}
