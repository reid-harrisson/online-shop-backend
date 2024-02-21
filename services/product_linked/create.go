package linkedsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(productID uint64, linkedID uint64, isUpCross uint64, modelProductLinked *models.ProductLinked) error {
	return service.DB.
		Create(&models.ProductLinked{
			ProductID: productID,
			LinkedID:  linkedID,
			IsUpCross: models.IsUpCross(isUpCross),
		}).
		Error
}
