package models

import (
	"OnlineStoreBackend/pkgs/utils"
	"time"
)

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

type StockAnalytics struct {
	Date     time.Time `gorm:"column:date"`
	StockIn  float64   `gorm:"column:stock_in"`
	StockOut float64   `gorm:"column:stock_out"`
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

type RepeatCustomerRates struct {
	ProductID  uint64 `gorm:"column:product_id"`
	CustomerID uint64 `gorm:"column:customer_id"`
}

type CustomerChurnRates struct {
	Rate float64 `gorm:"column:rate"`
}

type FullFunnelAnalytics struct {
	Page     utils.PageTypes `gorm:"column:page"`
	PageView uint64          `gorm:"column:page_views"`
}

type TopSellingProducts struct {
	ProductID   uint64  `gorm:"column:product_id"`
	ProductName string  `gorm:"column:product_name"`
	Sales       float64 `gorm:"column:sales"`
}

type OrderTrendAnalytics struct {
	Date  time.Time `gorm:"column:date"`
	Count uint64    `gorm:"column:count"`
	Sales float64   `gorm:"column:sales"`
}

type CustomerDataByLocation struct {
	Location  string `gorm:"column:location"`
	Customers uint64 `gorm:"column:customers"`
}

type CustomerSatisfaction struct {
	ProductID     uint64  `gorm:"column:product_id"`
	AverageRating float64 `gorm:"column:average_rating"`
	NPS           float64 `gorm:"column:nps"`
	Rating1       float64 `gorm:"column:rating1"`
	Rating2       float64 `gorm:"column:rating2"`
	Rating3       float64 `gorm:"column:rating3"`
	Rating4       float64 `gorm:"column:rating4"`
	Rating5       float64 `gorm:"column:rating5"`
}

type PageLoadingTime struct {
	Page        utils.PageTypes `gorm:"column:page"`
	AverageTime float64         `json:"average_time"`
	MaximumTime float64         `json:"maximum_time"`
	MinimumTime float64         `json:"minimum_time"`
}
