package models

import (
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model

	StoreID           uint64  `gorm:"column:store_id; type:bigint(20) unsigned"`
	Title             string  `gorm:"column:title; type:varchar(100)"`
	ShortDescription  string  `gorm:"column:short_description; type:varchar(100)"`
	LongDescription   string  `gorm:"column:long_description; type:varchar(500)"`
	ImageUrls         string  `gorm:"column:image_urls; type:varchar(1000)"`
	SKU               string  `gorm:"column:sku; type:varchar(45)"`
	UnitPriceRegular  float64 `gorm:"column:unit_price_regular; type:decimal(20,6)"`
	UnitPriceSale     float64 `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	StockQuantity     float64 `gorm:"column:stock_quantity; type:decimal(20,6)"`
	MinimumStockLevel float64 `gorm:"column:minimum_stock_level; type:decimal(20,6)"`
	CategoryID        uint64  `gorm:"column:category_id; type:bigint(20) unsigned"`
	Active            int8    `gorm:"column:active; type:tinyint(4)"`
}

type ProductsWithDetail struct {
	ProductsWithCategory
	RelatedChannels []ProductChannelsWithName   `gorm:"column:related_channels"`
	RelatedContents []ProductContentsWithTitle  `gorm:"column:related_contents"`
	Tags            []ProductTagsWithName       `gorm:"column:tags"`
	Attributes      []ProductAttributes         `gorm:"column:attributes"`
	Variations      []ProductVariationsWithName `gorm:"column:variations"`
	ShippingData    ShippingData                `gorm:"column:shipping_data"`
}

type ProductsWithCategory struct {
	Products
	Category string `gorm:"column:category"`
}

func (Products) TableName() string {
	return "store_products"
}
