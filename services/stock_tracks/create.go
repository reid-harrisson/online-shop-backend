package stocksvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelStock *models.StockTracks, req *requests.RequestStockTrack) error {
	modelStock.ProductID = req.ProductID
	modelStock.VariationID = req.VariationID
	modelStock.Change = req.Change
	modelStock.Event = req.Event

	return service.DB.Create(&modelStock).Error
}
