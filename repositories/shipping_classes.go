package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryShippingClass struct {
	DB *gorm.DB
}

func NewRepositoryShippingClass(db *gorm.DB) *RepositoryShippingClass {
	return &RepositoryShippingClass{DB: db}
}

func (repository *RepositoryShippingClass) ReadByID(modelClass *models.ShippingClasses, classID uint64) error {
	return repository.DB.
		Where("id = ?", classID).
		First(modelClass).
		Error
}

func (repository *RepositoryShippingClass) ReadByIDs(modelClasses *[]models.ShippingClasses, classIDs []uint64) error {
	return repository.DB.
		Where("id In (?)", classIDs).
		Find(modelClasses).
		Error
}
