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

func (repository *RepositoryVariation) ReadVariationsInStore(modelVars *[]models.ProductVariationsInStore, storeID uint64) {
	repository.DB.Table("store_product_variations As vars").
		Select(`
			vars.*,
			prods.title,
			prods.minimum_stock_level
		`).
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Where("vars.deleted_at Is Null And prods.deleted_at Is Null").
		Scan(&modelVars)
}

func (repository *RepositoryVariation) ReadVariationsInProduct(modelVars *[]models.ProductVariationsInProduct, productID uint64) {
	repository.DB.Table("store_product_variations As vars").
		Select(`
			vars.*,
			vals.attribute_id,
			vals.attribute_value,
			attrs.attribute_name,
			attrs.unit
		`).
		Joins("Left Join store_product_variation_details As dets On dets.variation_id = vars.id").
		Joins("Left Join store_product_attribute_values As vals On vals.id = dets.attribute_value_id").
		Joins("Left Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("vars.deleted_at Is Null And dets.deleted_at Is Null And vals.deleted_at Is Null And attrs.deleted_at Is Null").
		Scan(&modelVars)
}
