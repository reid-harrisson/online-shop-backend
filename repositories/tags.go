package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryTag struct {
	DB *gorm.DB
}

func NewRepositoryTag(db *gorm.DB) *RepositoryTag {
	return &RepositoryTag{DB: db}
}

func (repository *RepositoryTag) ReadByProductID(modelTags *[]models.ProductTagsWithName, productID uint64) error {
	return repository.DB.Table("store_product_tags As prodtags").
		Select("prodtags.*, tags.name As tag_name").
		Joins("Join store_tags As tags On tags.id = prodtags.tag_id").
		Where("tags.deleted_at Is Null And prodtags.deleted_at Is Null").
		Where("prodtags.product_id = ?", productID).
		Scan(modelTags).
		Error
}

func (repository *RepositoryTag) ReadByName(modelTag *models.Tags, name string, storeID uint64) error {
	return repository.DB.Where("name = ? And store_id = ?", name, storeID).First(modelTag).Error
}

func (repository *RepositoryTag) ReadByStoreID(modelTags *[]models.Tags, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).Find(modelTags).Error
}

func (repository *RepositoryTag) ReadByID(modelTag *models.Tags, tagID uint64) error {
	return repository.DB.Where("id = ?", tagID).First(modelTag).Error
}
