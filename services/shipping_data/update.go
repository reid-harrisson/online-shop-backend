package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(req *requests.RequestShippingMethod, modelShipData *models.ShippingData) error {
	service.DB.Where("product_id = ?", req.ProductID).First(modelShipData)
	service.DB.Save(&modelShipData)
	return nil
}
