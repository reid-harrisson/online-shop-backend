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
	return repository.DB.
		Where("store_id = ?", storeID).
		Find(modelMethods).
		Error
}

func (repository *RepositoryShippingMethod) ReadTableRateMethodByStoreID(modelMethod *models.ShippingMethods, storeID uint64) error {
	return repository.DB.
		Where("store_id = ?", storeID).
		First(modelMethod).
		Error
}

func (repository *RepositoryShippingMethod) ReadByID(modelMethod *models.ShippingMethods, methodID uint64) error {
	return repository.DB.
		Where("id = ?", methodID).
		First(modelMethod).
		Error
}

func (repository *RepositoryShippingMethod) ReadFlatRateByID(modelMethod *models.ShippingMethods, modelRates *[]models.ShippingFlatRates, methodID uint64) error {
	if err := repository.DB.Where("id = ?", methodID).First(modelMethod).Error; err != nil {
		return err
	}

	return repository.DB.
		Where("method_id = ?", methodID).
		Find(modelRates).
		Error
}

func (repository *RepositoryShippingMethod) ReadTableRateByID(modelMethod *models.ShippingMethods, modelRates *[]models.ShippingTableRates, methodID uint64) error {
	if err := repository.DB.Where("id = ?", methodID).First(modelMethod).Error; err != nil {
		return err
	}

	return repository.DB.
		Where("method_id = ?", methodID).
		Find(modelRates).
		Error
}

func (repository *RepositoryShippingMethod) ReadRates(modelRates *[]models.ShippingTableRates, storeID uint64) error {
	return repository.DB.
		Table("store_shipping_table_rates AS tables").
		Joins("Join store_shipping_methods AS methods On methods.id = tables.method_id").
		Where("methods.store_id = ? AND tables.deleted_at Is Null AND methods.deleted_at Is Null", storeID).
		Find(modelRates).
		Error
}

func (repository *RepositoryShippingMethod) ReadMethodAndTableRatesByStoreIDs(mapRates *map[uint64][]models.ShippingTableRates, mapMeth *map[uint64]models.ShippingMethods, storeIDs []uint64) error {
	modelMeths := []models.ShippingMethods{}
	modelRates := []models.ShippingTableRates{}

	if err := repository.DB.Where("store_id In (?)", storeIDs).Find(&modelMeths).Error; err != nil {
		return err
	}

	if err := repository.DB.Table("store_shipping_table_rates As tables").
		Joins("Join store_shipping_methods As methods On methods.id = tables.method_id").
		Where("methods.store_id In (?) And tables.deleted_at Is Null And methods.deleted_at Is Null", storeIDs).
		Find(&modelRates).Error; err != nil {
		return err
	}

	for _, modelMeth := range modelMeths {
		(*mapMeth)[modelMeth.StoreID] = modelMeth
	}

	for _, modelRate := range modelRates {
		(*mapRates)[modelRate.MethodID] = append((*mapRates)[modelRate.MethodID], modelRate)
	}

	return nil
}
