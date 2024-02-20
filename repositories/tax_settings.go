package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryTax struct {
	DB *gorm.DB
}

func NewRepositoryTax(db *gorm.DB) *RepositoryTax {
	return &RepositoryTax{DB: db}
}

func (repository *RepositoryTax) ReadTaxSetting(modelTax *models.TaxSettings, customerID uint64) error {
	return repository.DB.Table("users").
		Select(`users.country_id As country_id,
			countries.tax_rate As tax_rate,
			? As customer_id`, customerID).
		Joins("Join countries On countries.id = users.country_id").
		Where("users.id = ?", customerID).
		Where("countries.deleted_at Is Null And users.deleted_at Is Null").
		Limit(1).
		Scan(modelTax).
		Error
}
