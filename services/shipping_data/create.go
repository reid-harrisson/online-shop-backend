package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	service.DB.Where("product_id = ?", productID).Delete(models.ShippingData{})
	modelShipData.Classification = req.Classification
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Depth = req.Depth
	modelShipData.ProductID = productID
	service.DB.Create(&modelShipData)
	return nil
}
