package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryReview struct {
	DB *gorm.DB
}

func NewRepositoryReview(db *gorm.DB) *RepositoryReview {
	return &RepositoryReview{DB: db}
}

func (repository *RepositoryReview) ReadRate(modelProdRate *models.ProductRates, productID uint64) error {
	modelRates := make([]models.ProductCustomerRates, 0)
	repository.DB.Where("product_id = ?", productID).Find(&modelRates)
	modelProdRate.Customers = 0
	modelProdRate.Rate = 0.0
	for _, modelRate := range modelRates {
		modelProdRate.Customers++
		modelProdRate.Rate += modelRate.Rate
	}
	if modelProdRate.Customers != 0 {
		modelProdRate.Rate /= float64(modelProdRate.Customers)
	}
	modelProdRate.ProductID = productID
	return nil
}

func (repository *RepositoryReview) ReadReviews(modelReviews *[]models.ProductReviews, productID uint64) error {
	return repository.DB.
		Where("product_id = ?", productID).
		Find(modelReviews).
		Error
}

func (repository *RepositoryReview) ReadPublishReviews(modelReviews *[]models.ProductReviews, productID uint64) error {
	return repository.DB.
		Where("product_id = ? And status = ?", productID, 1).
		Find(modelReviews).
		Error
}
