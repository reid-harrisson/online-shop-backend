package repositories

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type RepositoryAnalytics struct {
	DB *gorm.DB
}

func NewRepositoryAnalytics(db *gorm.DB) *RepositoryAnalytics {
	return &RepositoryAnalytics{DB: db}
}

func (repository *RepositoryAnalytics) ReadSalesReport(modelReports *[]models.SalesReports, storeID uint64) error {
	return repository.DB.Table("store_order_items As items").
		Where("items.store_id = ?", storeID).
		Select(`
			items.variation_id,
			items.store_id,
			vars.product_id,
			items.price,
			items.quantity,
			items.sub_total_price As total_price
		`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Where("vars.deleted_at Is Null And items.deleted_at Is Null").
		Scan(modelReports).Error
}

func (repository *RepositoryAnalytics) ReadCustomerInsights(modelReport *models.CustomerInsights) error {
	return repository.DB.Table("users").
		Select(`
		  Min(users.age) As youngest_age,
    	Max(users.age) As oldest_age,
			Count(Distinct Case When Lower(users.gender) = 'male' Then users.id End) As male_count,
			Count(Distinct Case When Lower(users.gender) = 'female' Then users.id End) As female_count,
    	Avg(Distinct users.age) AS average_age
		`).
		Joins("Left Join store_orders As ords On ords.customer_id = users.id").
		Where("users.deleted_at Is Null And ords.deleted_at Is Null").
		Scan(modelReport).Error
}

func (repository *RepositoryAnalytics) ReadStockLevelAnalytics(modelLevels *[]models.StockLevelAnalytics, storeID uint64) error {
	return repository.DB.Table("store_product_variations As vars").
		Select(`
			Sum(vars.stock_level) As stock_level,
			vars.product_id,
      Case When Sum(vars.stock_level) > prods.minimum_stock_level Then 'Available' Else 'Out of Stock' End As availability
		`).
		Group("vars.product_id").
		Joins("Left Join store_products As prods On vars.product_id = prods.id").
		Where("vars.deleted_at Is Null And prods.deleted_at Is Null").
		Scan(modelLevels).Error
}

func (repository *RepositoryAnalytics) ReadVisitor(modelVisitor *models.VisitorAnalytics, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			Sum(Case When bounce = 1 Then 1 Else 0 End) As visitors,
			Count(Distinct ip_address) As unique_visitors,
			Count(id) As page_views,
			1 - Sum(Case When bounce = 2 Then 1 Else 0 End) / Sum(Case When bounce = 1 Then 1 Else 0 End) As bounce_rate
		`).
		Where("store_id = ?", storeID).
		Where("created_at > ? And created_at < ?", startDate, endDate).
		Scan(modelVisitor).Error
}

func (repository *RepositoryAnalytics) ReadConventionRate(modelRate *models.ConventionRate, storeID uint64) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			Sum(Case When page = ? Or page = ? Then 1 Else 0 End) / Sum(Case When bounce = 1 Then 1 Else 0 End) As rate
		`, utils.RegisterPage, utils.PaymentConfirmPage).
		Where("store_id = ?", storeID).
		Scan(modelRate).Error
}

func (repository *RepositoryAnalytics) ReadShoppingCartAbandonment(modelRate *models.ShoppingCartAbandonment, storeID uint64) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			Sum(Case When page = ? Then 1 Else 0 End) / Sum(Case When page = ? Then 1 Else 0 End) As rate
		`, utils.PaymentPage, utils.CartPage).
		Where("store_id = ?", storeID).
		Scan(modelRate).Error
}

func (repository *RepositoryAnalytics) ReadCheckoutFunnelAnalytics(modelSteps *[]models.CheckoutFunnelAnalytics, storeID uint64) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			page,
			Count(id) As page_views
		`).
		Group("page").
		Where("store_id = ?", storeID).
		Where("page In (?)", []utils.PageTypes{utils.CartPage, utils.PaymentPage, utils.PaymentConfirmPage}).
		Scan(modelSteps).Error
}

func (repository *RepositoryAnalytics) ReadFullFunnelAnalytics(modelSteps *[]models.FullFunnelAnalytics, storeID uint64) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			page,
			Count(id) As page_views
		`).
		Group("page").
		Where("store_id = ?", storeID).
		Scan(modelSteps).Error
}

func (repository *RepositoryAnalytics) ReadProductViewAnalytics(modelViews *[]models.ProductViewAnalytics, storeID uint64) error {
	return repository.DB.Model(models.Visitors{}).
		Select(`
			page,
			Count(id) As page_views
		`).
		Group("page").
		Scan(modelViews).Error
}

func (repository *RepositoryAnalytics) ReadRepeatCustomerRate(modelRate *[]models.RepeatCustomerRates, storeID uint64) error {
	return repository.DB.Table("store_order_items As items").
		Select(`
			vars.product_id,
			ords.customer_id
		`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Joins("Left Join store_orders As ords On ords.id = items.order_id").
		Group("ords.customer_id, items.order_id").
		Where("items.deleted_at Is Null And ords.deleted_at Is Null And vars.deleted_at Is Null").
		Scan(modelRate).
		Error
}

func (repository *RepositoryAnalytics) ReadCustomerChurnRate(modelRate *models.CustomerChurnRates, storeID uint64, startDate time.Time, endDate time.Time) error {
	activeUser := uint64(0)
	churnUser := 0
	err := repository.DB.Model(models.Orders{}).
		Select(`
			Distinct customer_id
		`).
		Count(&activeUser).
		Where("created_at > ? And created_at < ?", startDate, endDate).
		Count(&churnUser).
		Error
	modelRate.Rate = float64(churnUser) / float64(activeUser)
	return err
}

func (repository *RepositoryAnalytics) ReadTopSellingProducts(modelProducts *[]models.TopSellingProducts, storeID uint64, startDate time.Time, endDate time.Time, count uint64) error {
	return repository.DB.Table("store_order_items As items").
		Select(`
			prods.id As product_id,
			prods.title As product_name,
			Sum(items.quantity) As sales
		`).
		Joins("Join store_product_variations As vars On vars.id = items.variation_id").
		Joins("Join store_products As prods On prods.id = vars.product_id").
		Group("prods.id").
		Order("sales Desc").
		Where("prods.store_id = ?", storeID).
		Where("items.created_at > ? And items.created_at < ?", startDate, endDate).
		Limit(count).
		Scan(&modelProducts).Error
}

func (repository *RepositoryAnalytics) ReadOrderTrendAnalytics(modelTrends *[]models.OrderTrendAnalytics, storeID uint64, startDate time.Time, endDate time.Time) error {
	return repository.DB.Model(models.OrderItems{}).
		Select(`
			Date(created_at) As date,
			Count(Distinct order_id) As count,
			Sum(total_price) As sales
		`).
		Where("created_at Between ? And ?", startDate, endDate).
		Group("Date(created_at)").
		Scan(&modelTrends).Error
}
