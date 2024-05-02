package cousvc

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

func (service *Service) Delete(couponID uint64) error {
	query := service.DB.Delete(&models.Coupons{
		Model: gorm.Model{
			ID: uint(couponID),
		},
	})
	if query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
