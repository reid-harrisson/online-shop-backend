package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model

	StoreID          uint64              `gorm:"column:store_id; type:bigint(20) unsigned"`
	Title            string              `gorm:"column:title; type:varchar(100)"`
	ShortDescription string              `gorm:"column:short_description; type:varchar(100)"`
	LongDescription  string              `gorm:"column:long_description; type:varchar(500)"`
	ImageUrls        string              `gorm:"column:image_urls; type:varchar(1000)"`
	Status           utils.ProductStatus `gorm:"column:active; type:tinyint(4)"`
}

type ProductsWithDetail struct {
	Products

	RelatedChannels []ProductChannelsWithName   `gorm:"column:related_channels"`
	RelatedContents []ProductContentsWithTitle  `gorm:"column:related_contents"`
	Tags            []ProductTagsWithName       `gorm:"column:tags"`
	Categories      []ProductCategoriesWithName `gorm:"column:categories"`
	Attributes      []ProductAttributes         `gorm:"column:attributes"`
	Variations      []ProductVariationsWithName `gorm:"column:variations"`
	ShippingData    ShippingData                `gorm:"column:shipping_data"`
}

func (Products) TableName() string {
	return "store_products"
}
