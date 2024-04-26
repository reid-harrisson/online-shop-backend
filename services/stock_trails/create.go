package stocksvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelStock *models.StockTrails) error {
	return service.DB.Create(modelStock).Error
}

func (service *Service) CreateStocks(modelStocks *[]models.StockTrails) error {
	return service.DB.Create(modelStocks).Error
}
