package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryProductAttributeValue struct {
	DB *gorm.DB
}

func NewRepositoryProductAttributeValue(db *gorm.DB) *RepositoryProductAttributeValue {
	return &RepositoryProductAttributeValue{DB: db}
}

func (repository *RepositoryProductAttributeValue) ReadByID(modelVars *[]models.ProductAttributeValuesWithDetail, attributeID uint64) {
	repository.DB.Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name, attrs.unit As attribute_unit, attrs.product_id As product_id").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.attribute_id = ?", attributeID).
		Scan(modelVars)
}

func (repository *RepositoryProductAttributeValue) ReadByProductID(modelVars *[]models.ProductAttributeValuesWithDetail, productID uint64) {
	repository.DB.Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name, attrs.unit As attribute_unit, attrs.product_id As product_id").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("attrs.product_id = ?", productID).
		Order("vals.attribute_id").
		Scan(modelVars)
}

func (repository *RepositoryProductAttributeValue) ReadByIDs(modelValues *[]models.ProductAttributeValuesWithDetail, attributeValueIDs []uint64) {
	repository.DB.Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name, attrs.unit As attribute_unit, attrs.product_id As product_id").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.id In (?)", attributeValueIDs).
		Scan(modelValues)
}
