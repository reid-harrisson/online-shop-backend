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

func (repository *RepositoryVariation) ReadAllVariations(modelVars *[]models.ProductVariationsWithDetail) {
	repository.DB.Table("store_product_variations As vars").
		Select(`
			vars.*,
			prods.title,
			prods.minimum_stock_level
		`).
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Scan(&modelVars)
}
