package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryProductContent struct {
	DB *gorm.DB
}

func NewRepositoryProductContent(db *gorm.DB) *RepositoryProductContent {
	return &RepositoryProductContent{DB: db}
}

func (repository *RepositoryProductContent) ReadByProductID(modelContents *[]models.ProductContentsWithTitle, productID uint64) {
	repository.DB.Table("store_product_related_contents As prodconts").
		Select("prodconts.*, conts.title As content_title").
		Joins("Join contents As conts On conts.id = prodconts.content_id").
		Where("conts.deleted_at Is Null And prodconts.deleted_at Is Null").
		Where("prodconts.product_id = ?", productID).
		Scan(modelContents)
}
