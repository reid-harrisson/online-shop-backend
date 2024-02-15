package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryAttribute struct {
	DB *gorm.DB
}

func NewRepositoryAttribute(db *gorm.DB) *RepositoryAttribute {
	return &RepositoryAttribute{DB: db}
}

func (repository *RepositoryAttribute) ReadByProductID(modelAttrs *[]models.ProductAttributesWithName, productID uint64) {
	repository.DB.Table("store_product_attributes As prodattrs").
		Select("prodattrs.*, attrs.name As attribute_name, attrs.unit As attribute_unit").
		Joins("Join store_attributes As attrs On attrs.id = prodattrs.attribute_id").
		Where("attrs.deleted_at Is Null And prodattrs.deleted_at Is Null").
		Where("prodattrs.product_id = ?", productID).
		Scan(modelAttrs)
}
