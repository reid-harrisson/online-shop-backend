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

func (repository *RepositoryOrder) ReadByStoreID(modelOrders *[]models.ProductOrders, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).Find(modelOrders).Error
}

func (repository *RepositoryOrder) Read(modelOrders *[]models.ProductOrders, customerID uint64, productID uint64, storeID uint64) error {
	return repository.DB.
		Where("customer_id = ? Or ? = 0", customerID, storeID).
		Where("product_id = ? Or ? = 0", productID, storeID).
		Where("store_id = ? Or ? = 0", storeID, storeID).
		Find(modelOrders).Error
}

func (repository *RepositoryOrder) ReadByID(modelOrders *[]models.ProductOrders, id uint64) error {
	return repository.DB.Where("id = ?", id).Find(modelOrders).Error
}
