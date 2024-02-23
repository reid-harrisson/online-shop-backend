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

func (repository *RepositoryCart) ReadByID(modelItem *models.CartItems, cartID uint64) {
	repository.DB.First(modelItem, cartID)
}

func (repository *RepositoryCart) ReadByInfo(modelItem *models.CartItems, productID uint64, customerID uint64) {
	repository.DB.Where("customer_id = ? And product_id = ?", customerID, productID).First(modelItem)
}

func (repository *RepositoryCart) ReadByCustomerID(modelItems *[]models.CartItems, customerID uint64) {
	repository.DB.Where("customer_id = ?", customerID).Find(modelItems)
}

func (repository *RepositoryCart) ReadItemCount(modelCount *models.CartItemCount, customerID uint64) {
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
			prods.store_id As store_id,
			prods.unit_price_sale * carts.quantity As total_price, 
			prods.unit_price_sale As unit_price,
			prods.image_urls As image_url,
			prods.title As product_name,
			cates.name As category`).
		Joins("Left Join store_products As prods On prods.id = carts.product_id").
		Joins("Left Join store_product_categories As prodcates On prodcates.product_id = prods.id").
		Joins("Left Join store_categories As cates On cates.id = prodcates.category_id").
		Where("carts.customer_id = ?", customerID).
		Where("carts.deleted_at Is Null And prods.deleted_at Is Null And cates.deleted_at Is Null And prodcates.deleted_at Is Null").
		Scan(modelItems)
}
