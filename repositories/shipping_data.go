package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryShippingData struct {
	DB *gorm.DB
}

func NewRepositoryShippingData(db *gorm.DB) *RepositoryShippingData {
	return &RepositoryShippingData{DB: db}
}

func (repository *RepositoryShippingData) ReadByProductID(modelShipData *models.ShippingData, productID uint64) {
	repository.DB.Where("product_id = ?", productID).First(modelShipData)
}
