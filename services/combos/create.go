package combsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	coitmsvc "OnlineStoreBackend/services/combo_items"
	"encoding/json"
)

func (service *Service) Create(modelCombo *models.Combos, modelItems *[]models.ComboItems, req *requests.RequestCombo, storeID uint64) error {
	err := service.DB.Where("title = ?", req.Title).First(modelCombo).Error
	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelCombo.StoreID = storeID
	modelCombo.DiscountAmount = req.DiscountAmount
	modelCombo.DiscountType = utils.DiscountTypeFromString(req.DiscountType)
	modelCombo.ImageUrls = string(imageUrls)
	modelCombo.Description = req.Description
	modelCombo.Title = req.Title
	modelCombo.Status = utils.Draft
	if err == nil {
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
