package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryCategory struct {
	DB *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *RepositoryCategory {
	return &RepositoryCategory{DB: db}
}

func (repository *RepositoryCategory) ReadByProductID(modelCategories *[]models.ProductCategoriesWithName, productID uint64) {
	repository.DB.Table("store_product_categories As prodcates").
		Select("prodcates.*, cates.name As category_name").
		Joins("Join store_categories As cates On cates.id = prodcates.category_id").
		Where("cates.deleted_at Is Null And prodcates.deleted_at Is Null").
		Where("prodcates.product_id = ?", productID).
		Scan(modelCategories)
}
