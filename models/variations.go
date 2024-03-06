package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type ProductVariations struct {
	gorm.Model

	ProductID       uint64               `gorm:"column:product_id; type:bigint(20) unsigned"`
	Sku             string               `gorm:"column:sku; type:varchar(50)"`
	Price           float64              `gorm:"column:price; type:decimal(20,6)"`
	StockLevel      float64              `gorm:"column:stock_level; type:decimal(20,6)"`
	DiscountAmount  float64              `gorm:"column:discount_amount; type:decimal(20,6)"`
	DiscountType    utils.DiscountTypes  `gorm:"column:discount_type; type:tinyint(4)"`
	ImageUrls       string               `gorm:"column:image_urls; type:text"`
	Description     string               `gorm:"column:description; type:text"`
	Title           string               `gorm:"column:title; type:varchar(100)"`
	BackOrderStatus utils.SimpleStatuses `gorm:"column:back_order_status; type:tinyint(4)"`
}

func (ProductVariations) TableName() string {
	return "store_product_variations"
}

type ProductVariationsInStore struct {
	ProductVariations
	Title             string  `gorm:"column:title"`
	MinimumStockLevel float64 `gorm:"column:minimum_stock_level"`
}

type ProductVariationsInProduct struct {
	ProductVariations
	AttributeValueID uint64 `gomr:"column:attribute_value_id"`
	AttributeName    string `gorm:"column:attribute_name"`
	AttributeValue   string `gorm:"column:attribute_value"`
	Unit             string `gorm:"column:unit"`
}
