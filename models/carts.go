package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type CartItems struct {
	gorm.Model

	CustomerID  uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	VariationID uint64  `gorm:"column:variation_id; type:bigint(20) unsigned"`
	Quantity    float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

type CartItemsWithDetail struct {
	CartItems

	StoreID        uint64              `gorm:"column:store_id"`
	VariationName  string              `gorm:"column:variation_name"`
	ImageUrls      string              `gorm:"column:image_urls"`
	Price          float64             `gorm:"column:price"`
	Categories     string              `gorm:"column:categories"`
	DiscountAmount float64             `gorm:"column:discount_amount"`
	DiscountType   utils.DiscountTypes `gorm:"column:discount_type"`
	StockLevel     float64             `gorm:"column:stock_level"`
	Weight         float64             `gorm:"column:weight"`
	Width          float64             `gorm:"column:width"`
	Height         float64             `gorm:"column:height"`
	Length         float64             `gorm:"column:length"`
}

func (CartItems) TableName() string {
	return "store_cart_items"
}
