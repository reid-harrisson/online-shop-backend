package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryShippingData struct {
	DB *gorm.DB
}

func NewRepositoryShippingData(db *gorm.DB) *RepositoryShippingData {
	return &RepositoryShippingData{DB: db}
}

func (repository *RepositoryShippingData) ReadByVariationID(modelShip *models.ShippingData, variationID uint64) error {
	return repository.DB.
		Where("variation_id = ?", variationID).
		First(modelShip).
		Error
}
