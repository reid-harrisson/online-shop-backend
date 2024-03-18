package coitmsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
)

func (service *Service) Create(modelNewItems *[]models.ComboItems, req []requests.RequestComboItem, comboID uint64) error {
	variationIDs := []uint64{}
	indices := map[uint64]int{}
	for index, item := range req {
		if indices[item.VariationID] == 0 {
			*modelNewItems = append(*modelNewItems, models.ComboItems{
				ComboID:     comboID,
				Quantity:    item.Quantity,
				VariationID: item.VariationID,
			})
			variationIDs = append(variationIDs, item.VariationID)
			indices[item.VariationID] = index + 1
		}
	}

	modelCurItems := []models.ComboItems{}
	service.DB.Where("variation_id In (?) And combo_id = ?", variationIDs, comboID).Find(&modelCurItems)
	service.DB.Where("variation_id Not In (?) And combo_id = ?", variationIDs, comboID).Delete(&models.ComboItems{})
	for _, modelItem := range modelCurItems {
		index := indices[modelItem.VariationID] - 1
		(*modelNewItems)[index].ID = modelItem.ID
	}
	return service.DB.Save(modelNewItems).Error
}
