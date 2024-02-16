package shipsvc

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(productID uint64) error {
	return service.DB.Where("product_id = ?", productID).Delete(models.ShippingData{}).Error
}
