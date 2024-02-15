package models

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
