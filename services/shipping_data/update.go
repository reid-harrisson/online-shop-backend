package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(variationID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Length = req.Length
	modelShipData.VariationID = variationID

	return service.DB.Save(&modelShipData).Error
}
