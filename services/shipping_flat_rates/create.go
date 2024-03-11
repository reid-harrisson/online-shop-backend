package flatsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(methodID uint64, req []requests.RequestFlatRate, modelRates *[]models.ShippingFlatRates) error {
	classIDs := []uint64{}
	indices := map[uint64]int{}
	for index, rate := range req {
		if indices[rate.ClassID] == 0 {
			*modelRates = append(*modelRates, models.ShippingFlatRates{
				MethodID:    methodID,
				ClassID:     rate.ClassID,
				BaseCost:    rate.BaseCost,
				CostPerItem: rate.CostPerItem,
				Percent:     rate.Percent,
				MinFee:      rate.MinFee,
				MaxFee:      rate.MaxFee,
			})
			classIDs = append(classIDs, rate.ClassID)
			indices[rate.ClassID] = index + 1
		}
	}

	modelNewRates := []models.ShippingFlatRates{}
	service.DB.Where("class_id In (?) And method_id = ?", classIDs, methodID).Find(&modelNewRates)
	service.DB.Where("class_id Not In (?) And method_id = ?", classIDs, methodID).Delete(&models.ShippingFlatRates{})
	for _, modelRate := range modelNewRates {
		index := indices[modelRate.ClassID] - 1
		(*modelRates)[index].ID = modelRate.ID
	}
	return service.DB.Save(&modelRates).Error
}
