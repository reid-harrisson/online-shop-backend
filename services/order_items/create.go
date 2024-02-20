package orditmsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Create(modelItems []*models.OrderItems) {
	for _, modelItem := range modelItems {
		service.DB.Create(modelItem)
	}
}
