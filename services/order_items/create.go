package orditmsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelItems *[]models.OrderItems) {
	service.DB.Create(modelItems)
}
