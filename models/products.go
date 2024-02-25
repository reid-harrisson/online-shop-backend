package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model

	StoreID           uint64              `gorm:"column:store_id; type:bigint(20) unsigned"`
	Title             string              `gorm:"column:title; type:varchar(100)"`
	ShortDescription  string              `gorm:"column:short_description; type:varchar(100)"`
	LongDescription   string              `gorm:"column:long_description; type:varchar(500)"`
	ImageUrls         string              `gorm:"column:image_urls; type:varchar(1000)"`
	MinimumStockLevel float64             `gorm:"column:minimum_stock_level; type;decimal(20,6)"`
	CurrencyID        uint64              `gorm:"column:currency_id; type:bigint(20) unsigned"`
	Status            utils.ProductStatus `gorm:"column:status; type:tinyint(4)"`
}

type ProductsWithDetail struct {
	Products
	RelatedChannels []ProductChannelsWithName          `gorm:"column:related_channels"`
	RelatedContents []ProductContentsWithTitle         `gorm:"column:related_contents"`
	Tags            []ProductTagsWithName              `gorm:"column:tags"`
	Categories      []ProductCategoriesWithName        `gorm:"categories"`
	Attributes      []ProductAttributes                `gorm:"column:attributes"`
	AttributeValues []ProductAttributeValuesWithDetail `gorm:"column:variations"`
	ShippingData    ShippingData                       `gorm:"column:shipping_data"`
}

func (Products) TableName() string {
	return "store_products"
}
