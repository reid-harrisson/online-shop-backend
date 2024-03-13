package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryCoupon struct {
	DB *gorm.DB
}

func NewRepositoryCoupon(db *gorm.DB) *RepositoryCoupon {
	return &RepositoryCoupon{DB: db}
}

func (repository *RepositoryCoupon) ReadByStoreID(modelCoupons *[]models.Coupons, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).Find(modelCoupons).Error
}

func (repository *RepositoryCoupon) ReadByID(modelCoupon *models.Coupons, couponID uint64) error {
	return repository.DB.First(modelCoupon, couponID).Error
}

func (repository *RepositoryCoupon) ReadByCode(modelCoupon *models.Coupons, code string) error {
	return repository.DB.Where("coupon_code = ?", code).First(modelCoupon).Error
}

func (repository *RepositoryCoupon) ReadByIDs(modelCoupon *[]models.Coupons, ids []uint64) error {
	return repository.DB.Where("id In (?)", ids).Find(modelCoupon).Error
}
