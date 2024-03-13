package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryEmailTemplate struct {
	DB *gorm.DB
}

func NewRepositoryEmailTemplate(db *gorm.DB) *RepositoryEmailTemplate {
	return &RepositoryEmailTemplate{DB: db}
}

func (repository *RepositoryEmailTemplate) ReadByStoreID(modelStore *[]models.EmailTemplates, storeID uint64) error {
	return repository.DB.
		Where("store_id = ?", storeID).
		Find(modelStore).
		Error
}
