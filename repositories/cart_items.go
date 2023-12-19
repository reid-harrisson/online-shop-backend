package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryCartItem struct {
	DB *gorm.DB
}

func NewRepositoryCartItem(db *gorm.DB) *RepositoryCartItem {
	return &RepositoryCartItem{DB: db}
}

func (repository *RepositoryCartItem) Read(modelDetails *[]models.CartItemDetails, userID uint64) {
	repository.DB.Model(models.CartItems{}).
		Select(`
			store_cart_items.*,
			store_products.name As product,
			store_products.unit_price_sale As product_price,
			customers.first_name As user,
			vendors.first_name As store,
			vendors.id As store_id,
			store_cart_items.quantity As quantity
		`).
		Joins("Join store_products On store_products.id = store_cart_items.store_product_id").
		Joins("Join users As customers On customers.id = store_cart_items.user_id").
		Joins("Join users As vendors On vendors.id = store_products.user_id").
		Where("store_cart_items.user_id = ?", userID).
		Order("vendors.id").
		Scan(modelDetails)
}
