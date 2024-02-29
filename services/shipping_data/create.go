package shipsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"strconv"
)

func (service *Service) Create(productID uint64, req *requests.RequestShippingData, modelShipData *models.ShippingData) error {
	service.DB.Where("variation_id = ?", productID).Delete(models.ShippingData{})
	modelShipData.Classification = req.Classification
	modelShipData.Weight = req.Weight
	modelShipData.Width = req.Width
	modelShipData.Height = req.Height
	modelShipData.Length = req.Depth
	modelShipData.VariationID = productID
	service.DB.Create(&modelShipData)
	return nil
}

func (service *Service) CreateWithCSV(variationID uint64, modelCsv *models.CSVs) error {
	modelShip := models.ShippingData{}
	service.DB.Where("variation_id = ?", variationID).First(&modelShip)
	modelShip.Classification = modelCsv.ShippingClass
	modelShip.Weight, _ = strconv.ParseFloat(modelCsv.Weight, 64)
	modelShip.Width, _ = strconv.ParseFloat(modelCsv.Width, 64)
	modelShip.Height, _ = strconv.ParseFloat(modelCsv.Height, 64)
	modelShip.Length, _ = strconv.ParseFloat(modelCsv.Length, 64)
	modelShip.VariationID = variationID
	service.DB.Save(&modelShip)
	return nil
}
