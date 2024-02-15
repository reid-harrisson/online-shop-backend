package prodattrsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(attributeID uint64) {
	service.DB.Delete("id = ?", attributeID).Delete(models.ProductAttributes{})
}
