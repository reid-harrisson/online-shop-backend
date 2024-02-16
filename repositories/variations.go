package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryVariation struct {
	DB *gorm.DB
}

func NewRepositoryProductVariation(db *gorm.DB) *RepositoryVariation {
	return &RepositoryVariation{DB: db}
}

func (repository *RepositoryVariation) ReadByID(modelVars *[]models.ProductVariationsWithName, attributeID uint64) {
	repository.DB.Table("store_product_variations As vars").
		Select("vars.*, attrs.name As attribute_name, attrs.unit As attribute_unit").
		Joins("Join store_product_attributes As attrs On attrs.id = vars.attribute_id").
		Where("attrs.deleted_at Is Null And vars.deleted_at Is Null").
		Where("vars.attribute_id = ?", attributeID).
		Scan(modelVars)
}

func (repository *RepositoryVariation) ReadByProductID(modelVars *[]models.ProductVariationsWithName, productID uint64) {
	repository.DB.Table("store_product_variations As vars").
		Select("vars.*, attrs.name As attribute_name, attrs.unit As attribute_unit").
		Joins("Join store_product_attributes As attrs On attrs.id = vars.attribute_id").
		Where("attrs.deleted_at Is Null And vars.deleted_at Is Null").
		Where("attrs.product_id = ?", productID).
		Order("vars.attribute_id").
		Scan(modelVars)
}
