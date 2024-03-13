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
	service.DB.Save(modelCombo)
	itemService := coitmsvc.NewServiceComboItem(service.DB)
	return itemService.Create(modelItems, req.Items, uint64(modelCombo.ID))
}
