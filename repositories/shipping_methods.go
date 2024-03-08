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

func (repository *RepositoryShippingMethod) ReadByStoreID(modelMethods *[]models.ShippingMethods, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).Find(modelMethods).Error
}

func (repository *RepositoryShippingMethod) ReadDefault(modelMethod *models.ShippingMethods, storeID uint64) error {
	return repository.DB.Where("store_id = ?", storeID).First(modelMethod).Error
}

func (repository *RepositoryShippingMethod) ReadByID(modelMethod *models.ShippingMethods, methodID uint64) error {
	return repository.DB.Where("id = ?", methodID).First(modelMethod).Error
}

func (repository *RepositoryShippingMethod) ReadFlatRateByID(modelMethod *models.ShippingMethods, modelRates *[]models.ShippingFlatRates, methodID uint64) error {
	if err := repository.DB.Where("id = ?", methodID).First(modelMethod).Error; err != nil {
		return err
	}
	return repository.DB.Where("method_id = ?", methodID).Find(modelRates).Error
}
