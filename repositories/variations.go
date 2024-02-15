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

func (repository *RepositoryVariation) ReadByProductID(modelVars *[]models.ProductVariationsWithName, productID uint64) {
	repository.DB.Table("store_product_variations As prodvars").
		Select("prodvars.*, attrs.name As attribute_name, attrs.unit As attribute_unit").
		Joins("Join store_attributes As attrs On attrs.id = prodvars.attribute_id").
		Where("attrs.deleted_at Is Null And prodvars.deleted_at Is Null").
		Where("prodvars.product_id = ?", productID).
		Scan(modelVars)
}
