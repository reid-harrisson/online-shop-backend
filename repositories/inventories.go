package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryInventory struct {
	DB *gorm.DB
}

func NewRepositoryInventory(db *gorm.DB) *RepositoryInventory {
	return &RepositoryInventory{DB: db}
}

func (repository *RepositoryInventory) ReadInventories(modelInventories *[]models.Inventories, storeID uint64) error {
	return repository.DB.Table("store_product_variations As vars").
		Select(`
			prods.id As product_id,
			prods.title As product_name,
			vars.id As variation_id,
			vars.title As variation_name,
			vars.stock_level,
			prods.minimum_stock_level
		`).
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Where("prods.store_id = ?", storeID).
		Scan(modelInventories).
		Error
}
