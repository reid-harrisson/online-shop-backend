package combsvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	coitmsvc "OnlineStoreBackend/services/combo_items"
	"encoding/json"
)

func (service *Service) Update(modelCombo *models.Combos, modelItems *[]models.ComboItems, req *requests.RequestCombo, storeID uint64, comboID uint64) error {
	if err := service.DB.Where("id = ? And store_id = ?", comboID, storeID).First(modelCombo).Error; err != nil {
		return err
	}

	imageUrls, _ := json.Marshal(req.ImageUrls)
	modelCombo.StoreID = storeID
	modelCombo.DiscountAmount = req.DiscountAmount
	modelCombo.DiscountType = utils.DiscountTypeFromString(req.DiscountType)
	modelCombo.ImageUrls = string(imageUrls)
	modelCombo.Description = req.Description
	modelCombo.Title = req.Title
	modelCombo.Status = utils.Draft

	if err := service.DB.Save(modelCombo).Error; err != nil {
		return err
	}

	itemService := coitmsvc.NewServiceComboItem(service.DB)
	return itemService.Create(modelItems, req.Items, uint64(modelCombo.ID))
}

func (service *Service) UpdateStatus(status utils.ProductStatus, comboID uint64) error {
	return service.DB.Model(&models.Combos{}).Where("id = ?", comboID).Update("status", status).Error
}
