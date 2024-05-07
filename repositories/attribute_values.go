package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryAttributeValue struct {
	DB *gorm.DB
}

func NewRepositoryAttributeValue(db *gorm.DB) *RepositoryAttributeValue {
	return &RepositoryAttributeValue{DB: db}
}

func (repository *RepositoryAttributeValue) ReadByAttrID(modelVars *[]models.AttributeValuesWithDetail, attributeID uint64) error {
	return repository.DB.
		Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.attribute_id = ?", attributeID).
		Scan(modelVars).
		Error
}

func (repository *RepositoryAttributeValue) ReadByAttrIDAndValue(modelVars *models.AttributeValuesWithDetail, attributeID uint64, attributeValue string) error {
	return repository.DB.
		Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.attribute_id = ? And vals.attribute_value = ?", attributeID, attributeValue).
		First(modelVars).
		Error
}

func (repository *RepositoryAttributeValue) ReadByAttrValID(modelValue *models.AttributeValuesWithDetail, attributeValueID uint64) error {
	return repository.DB.
		Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.id = ?", attributeValueID).
		Scan(modelValue).
		Error
}

func (repository *RepositoryAttributeValue) ReadByProductID(modelVars *[]models.AttributeValuesWithDetail, productID uint64) error {
	return repository.DB.
		Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name, attrs.product_id As product_id").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("attrs.product_id = ?", productID).
		Order("vals.attribute_id").
		Scan(modelVars).
		Error
}

func (repository *RepositoryAttributeValue) ReadByIDs(modelValues *[]models.AttributeValuesWithDetail, attributeValueIDs []uint64) error {
	return repository.DB.
		Table("store_product_attribute_values As vals").
		Select("vals.*, attrs.attribute_name As attribute_name").
		Joins("Join store_product_attributes As attrs On attrs.id = vals.attribute_id").
		Where("attrs.deleted_at Is Null And vals.deleted_at Is Null").
		Where("vals.id In (?)", attributeValueIDs).
		Scan(modelValues).
		Error
}
