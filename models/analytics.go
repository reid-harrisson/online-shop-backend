package models

import "OnlineStoreBackend/pkgs/utils"

type StoreSales struct {
	StoreID uint64  `gorm:"column:store_id; type:bigint(20) unsigned"`
	Price   float64 `gorm:"column:price; type:decimal(20,6)"`
}

type ProductSales struct {
	ProductID uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	Quantity  float64 `gorm:"column:quantity; type:decimal(20,6)"`
	Total     float64 `gorm:"column:total; type:decimal(20,6)"`
}

type CategorySales struct {
	Category string  `gorm:"column:category; type:varchar(20)"`
	Quantity float64 `gorm:"column:quantity; type:decimal(20,6)"`
	Total    float64 `gorm:"column:total; type:decimal(20,6)"`
}

type CustomerSales struct {
	CustomerID uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Quantity   float64 `gorm:"column:quantity; type:decimal(20,6)"`
	Total      float64 `gorm:"column:total; type:decimal(20,6)"`
}

type Sales struct {
	ProductID   uint64  `gorm:"column:product_id"`
	ProductName string  `gorm:"column:product_name"`
	Price       float64 `gorm:"column:price"`
	Quantity    float64 `gorm:"column:quantity"`
	TotalPrice  float64 `gorm:"column:total_price"`
}

type SalesReports struct {
	VariationID uint64  `gorm:"column:variation_id"`
	ProductID   uint64  `gorm:"column:product_id"`
	StoreID     uint64  `gorm:"column:store_id"`
	Price       float64 `gorm:"column:price"`
	Quantity    float64 `gorm:"column:quantity"`
	TotalPrice  float64 `gorm:"column:total_price"`
}

type CustomerInsights struct {
	MaleCount   uint64  `gorm:"column:male_count"`
	FemaleCount uint64  `gorm:"column:female_count"`
	AverageAge  float64 `gorm:"column:average_age"`
	YoungestAge uint64  `gorm:"column:youngest_age"`
	OldestAge   uint64  `gorm:"column:oldest_age"`
	Location    string  `gorm:"column:location"`
}

type StockLevelAnalytics struct {
	ProductID    uint64  `gorm:"column:product_id"`
	StockLevel   float64 `gorm:"column:stock_level"`
	Availability string  `gorm:"column:availability"`
}

type VisitorAnalytics struct {
	Visitor       uint64  `gorm:"column:visitors"`
	UniqueVisitor uint64  `gorm:"column:unique_visitors"`
	PageView      uint64  `gorm:"column:page_views"`
	BounceRate    float64 `gorm:"column:bounce_rate"`
}

type ConventionRate struct {
	Rate float64 `gorm:"column:rate"`
}

type ShoppingCartAbandonment struct {
	Rate float64 `gorm:"column:rate"`
}

type CheckoutFunnelAnalytics struct {
	Page     utils.PageTypes `gorm:"column:page"`
	PageView uint64          `gorm:"column:page_views"`
}

type ProductViewAnalytics struct {
	ProductID   uint64 `gorm:"column:product_id"`
	ProductName string `gorm:"column:product_name"`
	PageView    uint64 `gorm:"column:page_views"`
	Purchase    uint64 `gorm:"column:purchase"`
}

type RepeatCustomerRate struct {
	ProductID  uint64 `gorm:"column:product_id"`
	CustomerID uint64 `gorm:"column:customer_id"`
}

type CustomerChurnRate struct {
	Rate float64 `gorm:"column:rate"`
}

type FullFunnelAnalytics struct {
	Page     utils.PageTypes `gorm:"column:page"`
	PageView uint64          `gorm:"column:page_views"`
}
