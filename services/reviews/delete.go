package revsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(id uint64) error {
	if err := service.DB.First(models.Reviews{}, id).Error; err != nil {
		return err
	}
	return service.DB.Delete(&models.Reviews{}, id).Error
}
