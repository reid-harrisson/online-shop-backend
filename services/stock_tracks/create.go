package stocksvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelStock models.StockTracks) error {
	return service.DB.Create(&modelStock).Error
}

func (service *Service) CreateStocks(modelStocks *[]models.StockTracks) error {
	return service.DB.Create(modelStocks).Error
}
