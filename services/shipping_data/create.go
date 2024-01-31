package shipData

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	service.DB.Where("store_product_id = ?", productID).Delete(models.ShippingData{})
	modelShipData.Classification = req.Classification
	modelShipData.Dimension = req.Dimension
	modelShipData.Weight = req.Weight
	modelShipData.ProductID = productID
	service.DB.Create(&modelShipData)
	return nil
}
