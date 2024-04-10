package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryReview struct {
	DB *gorm.DB
}

func NewRepositoryReview(db *gorm.DB) *RepositoryReview {
	return &RepositoryReview{DB: db}
}

func (repository *RepositoryReview) ReadReviews(modelReviews *[]models.Reviews, productID uint64) error {
	return repository.DB.
		Where("product_id = ?", productID).
		Find(modelReviews).
		Error
}

func (repository *RepositoryReview) ReadPublishReviews(modelReviews *[]models.Reviews, productID uint64) error {
	return repository.DB.
		Where("product_id = ? And status = ?", productID, 1).
		Find(modelReviews).
		Error
}
