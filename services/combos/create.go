package combsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	coitmsvc "OnlineStoreBackend/services/combo_items"
)

func (service *Service) Create(modelCombo *models.Combos, modelItems *[]models.ComboItems, req *requests.RequestCombo, storeID uint64) error {
	if err := service.DB.Where("title = ?", req.Title).First(modelCombo).Error; err == nil {
		if err = service.DB.Save(modelCombo).Error; err != nil {
			return err
		}
	} else {
		if err = service.DB.Create(modelCombo).Error; err != nil {
			return err
		}
	}

	itemService := coitmsvc.NewServiceComboItem(service.DB)
	return itemService.Create(modelItems, req.Items, uint64(modelCombo.ID))
}
