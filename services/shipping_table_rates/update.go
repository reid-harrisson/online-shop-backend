package tablesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Update(rateID uint64, req *requests.RequestTableRate, modelRates *models.ShippingTableRates) error {
	if err := service.DB.Where("id = ?", rateID).Find(&modelRates).Error; err != nil {
		return err
	}
	modelRates.Condition = req.Condition
	modelRates.Min = req.Min
	modelRates.Max = req.Max
	modelRates.RowCost = req.RowCost
	modelRates.ItemCost = req.ItemCost
	modelRates.CostPerKg = req.CostPerKg
	modelRates.PercentCost = req.PercentCost
	return service.DB.Save(modelRates).Error
}
