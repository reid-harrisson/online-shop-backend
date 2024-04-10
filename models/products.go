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
	RelatedChannels []ProductChannelsWithName   `gorm:"column:related_channels"`
	RelatedContents []ProductContentsWithTitle  `gorm:"column:related_contents"`
	Tags            []ProductTagsWithName       `gorm:"column:tags"`
	Categories      []ProductCategoriesWithName `gorm:"categories"`
	Attributes      []Attributes                `gorm:"column:attributes"`
	AttributeValues []AttributeValuesWithDetail `gorm:"column:variations"`
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

func (model *Products) AfterDelete(db *gorm.DB) (err error) {
	var modelAttrs = []Attributes{}
	db.Where("product_id = ?", model.ID).Find(&modelAttrs)
	db.Delete(&modelAttrs)

	var modelCates = []ProductCategories{}
	db.Where("product_id = ?", model.ID).Find(&modelCates)
	db.Delete(&modelCates)

	var modelLinks = []Links{}
	db.Where("product_id = ? And link_id = ?", model.ID, model.ID).Find(&modelLinks)
	db.Delete(&modelLinks)

	var modelChans = []ProductChannels{}
	db.Where("product_id = ?", model.ID).Find(&modelChans)
	db.Delete(&modelChans)

	var modelConts = []ProductContents{}
	db.Where("product_id = ?", model.ID).Find(&modelConts)
	db.Delete(&modelConts)

	var modelReviews = []Reviews{}
	db.Where("product_id = ?", model.ID).Find(&modelReviews)
	db.Delete(&modelReviews)

	var modelTags = []ProductTags{}
	db.Where("product_id = ?", model.ID).Find(&modelTags)
	db.Delete(&modelTags)

	db.Where("badge_id = ?", model.ID).Delete("invoice_item")

	return
}
