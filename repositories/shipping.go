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

func (repository *RepositoryShipping) ReadByVariationID(modelShip *models.ShippingData, variationID uint64) {
	repository.DB.Where("variation_id = ?", variationID).First(modelShip)
}

func (repository *RepositoryShipping) ReadByStoreID(modelMethods *[]models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).Find(modelMethods)
}

func (repository *RepositoryShipping) ReadDefaultMethod(modelMethod *models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).First(modelMethod)
}

func (repository *RepositoryShipping) ReadByID(modelMethod *models.ShippingMethods, methodID uint64) {
	repository.DB.Where("id = ?", methodID).First(modelMethod)
}
