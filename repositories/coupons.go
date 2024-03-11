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

func (repository *RepositoryCoupon) ReadByStoreID(modelAddrs *[]models.Coupons, storeID uint64) {
	repository.DB.Where("store_id = ? And active = 1", storeID).Find(modelAddrs)
}
