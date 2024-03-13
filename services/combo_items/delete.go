package coitmsvc

import "OnlineStoreBackend/models"

func (service *Service) Delete(couponID uint64) error {
	return service.DB.Where("id = ?", couponID).Delete(&models.Coupons{}).Error
}
