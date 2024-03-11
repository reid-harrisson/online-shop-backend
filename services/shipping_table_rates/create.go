package tablesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Create(methodID uint64, req *requests.RequestTableRate, modelRate *models.ShippingTableRates) error {
	condition := utils.ConditionFromString(req.Condition)
	err := service.DB.Where("`condition` = ? And min = ? And max = ?", condition, req.Min, req.Max).First(&modelRate).Error
	modelRate.Condition = condition
	modelRate.Min = req.Min
	modelRate.Max = req.Max
	modelRate.RowCost = req.RowCost
	modelRate.ItemCost = req.ItemCost
	modelRate.CostPerKg = req.CostPerKg
	modelRate.PercentCost = req.PercentCost
	modelRate.MethodID = methodID
	if err != nil {
		return service.DB.Create(modelRate).Error
	}
	return nil
}

func (service *Service) CreateMany(methodID uint64, req []requests.RequestTableRate, modelRates *[]models.ShippingTableRates) error {
	matches := []string{}
	indices := map[string]int{}
	for index, rate := range req {
		condition := utils.ConditionFromString(rate.Condition)
		match := fmt.Sprintf("%d:%f:%f", condition, rate.Min, rate.Max)
		if indices[match] == 0 {
			*modelRates = append(*modelRates, models.ShippingTableRates{
				MethodID:    methodID,
				Condition:   condition,
				Min:         rate.Min,
				Max:         rate.Max,
				RowCost:     rate.RowCost,
				ItemCost:    rate.ItemCost,
				CostPerKg:   rate.CostPerKg,
				PercentCost: rate.PercentCost,
			})
			matches = append(matches, match)
			indices[match] = index + 1
		}
	}

	modelNewRates := []models.ShippingTableRates{}
	service.DB.Where(`Concat(class_id,':',"condition",':',min,':',max) In (?) And method_id = ?`, matches, methodID).Find(&modelNewRates)
	service.DB.Where(`Concat(class_id,':',"condition",':',min,':',max) Not In (?) And method_id = ?`, matches, methodID).Delete(&models.ShippingTableRates{})
	for _, modelRate := range modelNewRates {
		match := fmt.Sprint(modelRate.ClassID, ':', modelRate.Condition, ':', modelRate.Min, ':', modelRate.Max)
		index := indices[match] - 1
		(*modelRates)[index].ID = modelRate.ID
	}
	return service.DB.Save(&modelRates).Error
}
