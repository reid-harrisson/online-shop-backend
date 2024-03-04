package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositorySales struct {
	DB *gorm.DB
}

func NewRepositorySales(db *gorm.DB) *RepositorySales {
	return &RepositorySales{DB: db}
}

func (repository *RepositorySales) ReadSalesByProduct(modelSales *[]models.ProductSales, storeID uint64) error {
	return repository.DB.Model(models.Orders{}).Where("store_id = ?", storeID).
		Select(`
			product_id As product_id,
			Sum(quantity) As quantity,
			Sum(total_price) As total
		`).
		Group("product_id").Order("total Desc").Scan(modelSales).Error
}

func (repository *RepositorySales) ReadSalesByCategory(modelSales *[]models.CategorySales, storeID uint64) error {
	return repository.DB.Table("store_order_items As prodOdrs").
		Where("prodOdrs.store_id = ? And prodOdrs.deleted_at Is Null", storeID).
		Select(`
			prodTags.tag As category,
			Sum(prodOdrs.quantity) As quantity,
			Sum(prodOdrs.total_price) As total
		`).
		Joins("Right Join store_product_tags As prodTags On prodTags.deleted_at Is Null And prodOdrs.product_id = prodTags.product_id").
		Group("prodTags.tag").Order("total Desc").Scan(modelSales).Error
}

func (repository *RepositorySales) ReadCLV(modelSales *[]models.CustomerSales, storeID uint64) error {
	return repository.DB.Model(models.OrderItems{}).Where("store_id = ?", storeID).
		Select(`
			customer_id,
			Sum(quantity) As quantity,
			Sum(total_price) As total
		`).
		Group("customer_id").Order("total Desc").Scan(modelSales).Error
}

func (repository *RepositorySales) ReadRevenue(modelSale *models.StoreSales, storeID uint64) error {
	return repository.DB.Model(models.OrderItems{}).Where("store_id = ?", storeID).
		Select("store_id,	Sum(total_price) As price").Scan(modelSale).Error
}

func (repository *RepositorySales) ReadAOV(modelSale *models.StoreSales, storeID uint64) error {
	return repository.DB.Model(models.OrderItems{}).Where("store_id = ?", storeID).
		Select("store_id,	Avg(total_price) As price").Scan(modelSale).Error
}
