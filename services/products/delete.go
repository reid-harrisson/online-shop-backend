package product

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(productID uint64) error {
	return service.DB.Where("id = ?", productID).Delete(models.Products{}).Error
}
