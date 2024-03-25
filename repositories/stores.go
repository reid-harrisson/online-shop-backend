package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryStore struct {
	DB *gorm.DB
}

func NewRepositoryStore(db *gorm.DB) *RepositoryStore {
	return &RepositoryStore{DB: db}
}

func (repository *RepositoryStore) ReadByID(modelStore *models.Stores, storeID uint64) error {
	return repository.DB.First(modelStore, storeID).Error
}

func (repository *RepositoryStore) ReadAll(modelStores *[]models.Stores) error {
	return repository.DB.Find(modelStores).Error
}

func (repository *RepositoryStore) ReadByUser(modelStores *[]models.Stores, userID uint64) error {
	return repository.DB.Where("owner_id = ?", userID).Find(modelStores).Error
}
