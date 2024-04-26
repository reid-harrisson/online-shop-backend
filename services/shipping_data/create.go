package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(variationID uint64, req *requests.RequestShippingData, modelShip *models.ShippingData) error {
	modelShip.Weight = req.Weight
	modelShip.Width = req.Width
	modelShip.Height = req.Height
	modelShip.Length = req.Length
	modelShip.VariationID = variationID
	service.DB.Create(&modelShip)
	return nil
}

func (service *Service) CreateWithCSV(modelNewShips *[]models.ShippingData, shipVarIDs []uint64, shipIndices map[uint64]int) error {
	modelCurShips := []models.ShippingData{}
	if err := service.DB.Where("variation_id In (?)", shipVarIDs).Find(&modelCurShips).Error; err != nil {
		return err
	}
	for _, modelShip := range modelCurShips {
		index := shipIndices[modelShip.VariationID]
		(*modelNewShips)[index].ID = modelShip.ID
	}
	return service.DB.Save(modelNewShips).Error
}
