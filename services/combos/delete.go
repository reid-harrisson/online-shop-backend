package combsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(storeID uint64, comboID uint64) error {
	if err := service.DB.Where("combo_id = ?", comboID).Delete(&models.ComboItems{}).Error; err != nil {
		return err
	}
	return service.DB.Where("id = ? And store_id = ?", comboID, storeID).Delete(&models.Combos{}).Error
}
