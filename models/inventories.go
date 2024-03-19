package models

type Inventories struct {
	ProductID         uint64  `gorm:"column:product_id"`
	ProductName       string  `gorm:"column:product_name"`
	VariationID       uint64  `gorm:"column:variation_id"`
	VariationName     string  `gorm:"column:variation_name"`
	StockLevel        float64 `gorm:"column:stock_level"`
	MinimumStockLevel float64 `gorm:"column:minimum_stock_level"`
}
