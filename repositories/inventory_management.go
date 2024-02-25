package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryInventory struct {
	DB *gorm.DB
}

func NewRepositoryInventory(db *gorm.DB) *RepositoryInventory {
	return &RepositoryInventory{DB: db}
}

func (repository *RepositoryInventory) ReadOne(modelStore *models.Stores, storeID uint64) error {
	return repository.DB.
		Where("id = ?", storeID).
		Find(modelStore).
		Error
}
