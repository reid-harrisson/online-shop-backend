package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryVariation struct {
	DB *gorm.DB
}

func NewRepositoryVariation(db *gorm.DB) *RepositoryVariation {
	return &RepositoryVariation{DB: db}
}

func (repository *RepositoryVariation) ReadVariationByID(modelVar *models.ProductVariations, variationID uint64) {
	repository.DB.First(modelVar, variationID)
}

func (repository *RepositoryVariation) ReadAllVariations(modelVar *[]models.ProductVariations) {
	repository.DB.Find(modelVar)
}
