package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Length = req.Length
	modelShipData.VariationID = productID

	return service.DB.Save(&modelShipData).Error
}
