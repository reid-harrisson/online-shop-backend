package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryShipping struct {
	DB *gorm.DB
}

func NewRepositoryShipping(db *gorm.DB) *RepositoryShipping {
	return &RepositoryShipping{DB: db}
}

func (repository *RepositoryShipping) ReadByProductID(modelShipData *models.ShippingData, productID uint64) {
	repository.DB.Where("product_id = ?", productID).First(modelShipData)
}

func (repository *RepositoryShipping) ReadOptionsByStoreID(modelOptions *[]models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).Find(modelOptions)
}
