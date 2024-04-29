package tablesvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(rateID uint64) error {
	return service.DB.
		Where("id = ?", rateID).
		Delete(&models.ShippingTableRates{}).
		Error
}
