package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryShippingMethod struct {
	DB *gorm.DB
}

func NewRepositoryShippingMethod(db *gorm.DB) *RepositoryShippingMethod {
	return &RepositoryShippingMethod{DB: db}
}

func (repository *RepositoryShippingMethod) ReadByStoreID(modelMethods *[]models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).Find(modelMethods)
}

func (repository *RepositoryShippingMethod) ReadDefaultMethod(modelMethod *models.ShippingMethods, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).First(modelMethod)
}

func (repository *RepositoryShippingMethod) ReadByID(modelMethod *models.ShippingMethods, methodID uint64) {
	repository.DB.Where("id = ?", methodID).First(modelMethod)
}
