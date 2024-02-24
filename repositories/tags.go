package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryTag struct {
	DB *gorm.DB
}

func NewRepositoryTag(db *gorm.DB) *RepositoryTag {
	return &RepositoryTag{DB: db}
}

func (repository *RepositoryTag) ReadByProductID(modelTags *[]models.ProductTagsWithName, productID uint64) {
	repository.DB.Table("store_product_tags As prodtags").
		Select("prodtags.*, tags.name As tag_name").
		Joins("Join store_tags As tags On tags.id = prodtags.tag_id").
		Where("tags.deleted_at Is Null And prodtags.deleted_at Is Null").
		Where("prodtags.product_id = ?", productID).
		Scan(modelTags)
}

func (repository *RepositoryTag) ReadByName(modelTag *models.StoreTags, name string) {
	repository.DB.Where("name = ?", name).First(modelTag)
}

func (repository *RepositoryTag) ReadAll(modelTags *[]models.StoreTags) {
	repository.DB.Find(modelTags)
}
