package linksvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(linkedProductID uint64) {
	service.DB.
		Where("id = ?", linkedProductID).
		Delete(&models.ProductLinks{})
}
