package cart

import (
	"OnlineStoreBackend/models"
)

func (service *Service) Delete(cartID uint64) error {
	return service.DB.Delete(models.CartItems{}, cartID).Error
}

func (service *Service) DeleteAll(customerID uint64) error {
	return service.DB.Where("customer_id = ?", customerID).Delete(models.CartItems{}).Error
}
