package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryOrder struct {
	DB *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{DB: db}
}

func (repository *RepositoryOrder) Read(modelOrders *[]models.ProductOrders, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).Find(modelOrders).Error
}
