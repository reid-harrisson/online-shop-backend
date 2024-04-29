package linksvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(linkedProductID uint64) error {
	return service.DB.
		Where("id = ?", linkedProductID).
		Delete(&models.Links{}).
		Error
}
