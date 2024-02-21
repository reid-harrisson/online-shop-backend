package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryEmailTemplate struct {
	DB *gorm.DB
}

func NewRepositoryEmailTemplate(db *gorm.DB) *RepositoryEmailTemplate {
	return &RepositoryEmailTemplate{DB: db}
}

func (repository *RepositoryEmailTemplate) ReadEmailTemplateByStoreID(modelStore *[]models.EmailTemplate, storeID uint64) error {
	return repository.DB.
		Where("store_id = ?", storeID).
		Find(modelStore).
		Error
}
