package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryCart struct {
	DB *gorm.DB
}

func NewRepositoryCart(db *gorm.DB) *RepositoryCart {
	return &RepositoryCart{DB: db}
}

func (repository *RepositoryCart) ReadAll(modelCarts *[]models.CartItemWithPrice, customerID uint64, storeID uint64) error {
	return repository.DB.Model(models.CartItems{}).
		Select(`store_cart_items.*,
			store_products.unit_price_sale * store_cart_items.quantity As price, 
			store_products.unit_price_sale As unit_price_sale`).
		Joins("Left Join store_products On store_products.id = store_cart_items.product_id").
		Where("customer_id = ?", customerID).
		Where("? = 0 Or store_cart_items.store_id = ?", storeID, storeID).
		Scan(modelCarts).Error
}

func (repository *RepositoryCart) ReadPreview(modelCarts *[]models.CartItemWithPrice, modelTaxSet *models.TaxSettings, customerID uint64) error {
	if err := repository.DB.Model(models.CartItems{}).
		Select(`store_cart_items.*,
			store_products.unit_price_sale * store_cart_items.quantity As price, 
			store_products.unit_price_sale As unit_price_sale`).
		Joins("Left Join store_products On store_products.id = store_cart_items.product_id").
		Where("customer_id = ?", customerID).
		Scan(modelCarts).Error; err != nil {
		return err
	}
	return repository.DB.Table("users").
		Select(`users.country_id As country_id,
			countries.tax_rate As tax_rate`).
		Joins("Join countries On countries.id = users.country_id").
		Where("users.id = ?", customerID).
		Limit(1).
		Scan(modelTaxSet).
		Error
}
