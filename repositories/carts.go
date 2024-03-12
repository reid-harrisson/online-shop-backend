package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryCart struct {
	DB *gorm.DB
}

func NewRepositoryCart(db *gorm.DB) *RepositoryCart {
	return &RepositoryCart{DB: db}
}

func (repository *RepositoryCart) ReadByID(modelItem *models.CartItems, cartID uint64) {
	repository.DB.First(modelItem, cartID)
}

func (repository *RepositoryCart) ReadByInfo(modelItem *models.CartItems, variationID uint64, customerID uint64) {
	repository.DB.Where("customer_id = ? And variation_id = ?", customerID, variationID).First(modelItem)
}

func (repository *RepositoryCart) ReadByCustomerID(modelItems *[]models.CartItems, customerID uint64) {
	repository.DB.Where("customer_id = ?", customerID).Find(modelItems)
}

func (repository *RepositoryCart) ReadItemCount(modelCount *models.CartCount, customerID uint64) {
	repository.DB.
		Model(models.CartItems{}).
		Select(`
			Count(id) As count
		`).
		Where("customer_id = ?", customerID).
		Scan(modelCount)
}

func (repository *RepositoryCart) ReadDetail(modelItems *[]models.CartItemsWithDetail, customerID uint64) {
	repository.DB.
		Table("store_cart_items As carts").
		Select(`carts.*,
			prods.store_id,
			vars.price,
			vars.discount_amount,
			vars.discount_type,
			vars.image_urls,
			vars.stock_level,
			vars.title As variation_name,
			Group_Concat(Concat('"', cates.name,'"') Separator ', ') As categories`).
		Joins("Left Join store_product_variations As vars On vars.id = carts.variation_id").
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Joins("Left Join store_product_categories As prodcates On prodcates.product_id = prods.id").
		Joins("Left Join store_categories As cates On cates.id = prodcates.category_id").
		Group("carts.id").
		Where("carts.customer_id = ?", customerID).
		Where("carts.deleted_at Is Null And prods.deleted_at Is Null And cates.deleted_at Is Null And prodcates.deleted_at Is Null").
		Scan(modelItems)
}
