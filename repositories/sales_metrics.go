package repositories

import (
	"OnlineStoreBackend/models"
	"time"

	"gorm.io/gorm"
)

type RepositorySales struct {
	DB *gorm.DB
}

func NewRepositorySales(db *gorm.DB) *RepositorySales {
	return &RepositorySales{DB: db}
}

func (repository *RepositorySales) ReadSalesByProduct(modelSales *[]models.ProductSales, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Table("store_order_items As items").
		Select(`
			vars.product_id As product_id,
			Sum(items.quantity) As quantity,
			Sum(items.total_price) As total
		`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Group("product_id").Order("total Desc").
		Where("items.store_id = ? And items.created_at Between ? And ?", storeID, startDate, endDate).
		Where("items.deleted_at Is Null And vars.deleted_at Is Null").
		Scan(modelSales).Error
}

func (repository *RepositorySales) ReadSalesByCategory(modelSales *[]models.CategorySales, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Table("store_order_items As items").
		Select(`
			cates.name As category,
			Sum(items.quantity) As quantity,
			Sum(items.total_price) As total
		`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Joins("Right Join store_product_categories As pcates On vars.product_id = pcates.product_id").
		Joins("Left Join store_categories As cates On cates.id = pcates.product_id").
		Group("cates.name").Order("total Desc").
		Where("items.store_id = ? And items.created_at Between ? And ?", storeID, startDate, endDate).
		Where("items.deleted_at Is Null And vars.deleted_at Is Null And cates.deleted_at Is Null And pcates.deleted_at Is Null").
		Scan(modelSales).Error
}

func (repository *RepositorySales) ReadCLV(modelSales *[]models.CustomerSales, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Table("store_order_items As items").
		Select(`
			ords.customer_id,
			Sum(items.quantity) As quantity,
			Sum(items.total_price) As total
		`).
		Joins("Left Join store_orders As ords On ords.id = items.order_id").
		Group("ords.customer_id").Order("total Desc").
		Where("items.store_id = ? And items.created_at Between ? And ?", storeID, startDate, endDate).
		Where("items.deleted_at Is Null And ords.deleted_at Is Null").
		Scan(modelSales).Error
}

func (repository *RepositorySales) ReadRevenue(modelSale *models.StoreSales, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Model(models.OrderItems{}).
		Select("store_id,	Sum(total_price) As price").
		Where("store_id = ? And created_at Between ? And ?", storeID, startDate, endDate).
		Scan(modelSale).Error
}

func (repository *RepositorySales) ReadAOV(modelSale *models.StoreSales, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Model(models.OrderItems{}).
		Select("store_id,	Avg(total_price) As price").
		Where("store_id = ? And created_at Between ? And ?", storeID, startDate, endDate).
		Scan(modelSale).Error
}
