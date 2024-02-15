package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) UpdateShippingData(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Depth = req.Depth
	modelShipData.ProductID = productID
	service.DB.Save(&modelShipData)
	return nil
}

func (service *Service) Update(req *requests.RequestShippingMethod, modelShipData *models.ShippingData) error {
	service.DB.Where("product_id = ?", req.ProductID).First(modelShipData)
	service.DB.Save(&modelShipData)
	return nil
}
