package tablesvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
)

func (service *Service) Create(methodID uint64, req []requests.RequestTableRate, modelRates *[]models.ShippingTableRates) error {
	matches := []string{}
	indices := map[string]int{}
	for index, rate := range req {
		match := fmt.Sprintf("%d:%d:%f:%f", rate.ClassID, rate.Condition, rate.Min, rate.Max)
		if indices[match] == 0 {
			*modelRates = append(*modelRates, models.ShippingTableRates{
				MethodID:    methodID,
				ClassID:     rate.ClassID,
				Condition:   rate.Condition,
				Min:         rate.Min,
				Max:         rate.Max,
				Break:       rate.Break,
				Abort:       rate.Abort,
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
