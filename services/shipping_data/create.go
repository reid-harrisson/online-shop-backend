package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	service.DB.Where("variation_id = ?", productID).Delete(models.ShippingData{})
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Length = req.Length
	modelShipData.VariationID = productID
	service.DB.Create(&modelShipData)
	return nil
}

func (service *Service) CreateWithCSV(modelNewShips *[]models.ShippingData, shipVarIDs []uint64, shipIndices map[uint64]int) {
	modelCurShips := []models.ShippingData{}
	service.DB.Where("variation_id In (?)", shipVarIDs).Find(&modelCurShips)
	for _, modelShip := range modelCurShips {
		index := shipIndices[modelShip.VariationID]
		(*modelNewShips)[index].ID = modelShip.ID
	}
	service.DB.Save(modelNewShips)
}
