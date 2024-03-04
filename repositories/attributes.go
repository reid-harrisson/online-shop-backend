package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryAttribute struct {
	DB *gorm.DB
}

func NewRepositoryAttribute(db *gorm.DB) *RepositoryAttribute {
	return &RepositoryAttribute{DB: db}
}

func (repository *RepositoryAttribute) ReadByProductID(modelAttrs *[]models.ProductAttributes, productID uint64) {
	repository.DB.Where("product_id = ?", productID).Find(modelAttrs)
}

func (repository *RepositoryAttribute) ReadByName(modelAttr *models.ProductAttributes, name string) {
	repository.DB.Where("name = ?", name).First(modelAttr)
}

func (repository *RepositoryAttribute) ReadByID(modelAttr *models.ProductAttributes, attributeID uint64) {
	repository.DB.Where("id = ?", attributeID).First(modelAttr)
}
