package orditmsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelItems *[]models.OrderItems) error {
	return service.DB.Create(modelItems).Error
}
