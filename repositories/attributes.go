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

func (repository *RepositoryAttribute) ReadByProductID(modelAttrs *[]models.Attributes, productID uint64) error {
	return repository.DB.Where("product_id = ?", productID).Find(modelAttrs).Error
}

func (repository *RepositoryAttribute) ReadByName(modelAttr *models.Attributes, name string) {
	repository.DB.Where("name = ?", name).First(modelAttr)
}

func (repository *RepositoryAttribute) ReadByID(modelAttr *models.Attributes, attributeID uint64) {
	repository.DB.Where("id = ?", attributeID).First(modelAttr)
}
