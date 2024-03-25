package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model

	StoreID           uint64              `gorm:"column:store_id; type:bigint(20) unsigned"`
	Title             string              `gorm:"column:title; type:varchar(100)"`
	ShortDescription  string              `gorm:"column:short_description; type:varchar(100)"`
	LongDescription   string              `gorm:"column:long_description; type:varchar(500)"`
	ImageUrls         string              `gorm:"column:image_urls; type:varchar(1000)"`
	MinimumStockLevel float64             `gorm:"column:minimum_stock_level; type;decimal(20,6)"`
	Status            utils.ProductStatus `gorm:"column:status; type:tinyint(4)"`
	Sku               string              `gorm:"column:sku; type:varchar(50)"`
	Type              utils.ProductTypes  `gorm:"column:type; type:tinyint(4)"`
	ShippingClass     string              `gorm:"column:shipping_class; type:varchar(100)"`
}

type ProductsWithDetail struct {
	Products
	RelatedChannels []ProductChannelsWithName          `gorm:"column:related_channels"`
	RelatedContents []ProductContentsWithTitle         `gorm:"column:related_contents"`
	Tags            []ProductTagsWithName              `gorm:"column:tags"`
	Categories      []ProductCategoriesWithName        `gorm:"categories"`
	Attributes      []ProductAttributes                `gorm:"column:attributes"`
	AttributeValues []ProductAttributeValuesWithDetail `gorm:"column:variations"`
}

type ProductsApproved struct {
	ID           uint64  `gorm:"column:id"`
	Title        string  `gorm:"column:title"`
	MinimumPrice float64 `gorm:"column:minimum_price"`
	MaximumPrice float64 `gorm:"column:maximum_price"`
	RegularPrice float64 `gorm:"column:regular_price"`
	Rating       float64 `gorm:"column:rating"`
	ImageUrls    string  `gorm:"column:image_urls"`
}

func (Products) TableName() string {
	return "store_products"
}
