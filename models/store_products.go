package models

import (
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model

	StoreID          uint64  `gorm:"column:store_id; type:bigint(20) unsigned"`
	Name             string  `gorm:"column:name; type:varchar(100)"`
	Brief            string  `gorm:"column:brief; type:varchar(100)"`
	Description      string  `gorm:"column:description; type:varchar(500)"`
	ImageUrls        string  `gorm:"column:image_urls; type:varchar(1000)"`
	SKU              string  `gorm:"column:sku; type:varchar(45)"`
	UnitPriceRegular float64 `gorm:"column:unit_price_regular; type:decimal(20,6)"`
	UnitPriceSale    float64 `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	StockQuantity    float64 `gorm:"column:stock_quantity; type:decimal(20,6)"`
	LinkedProductIDs string  `gorm:"column:linked_product_ids; type:varchar(100)"`
	Active           int8    `gorm:"column:active; type:tinyint(4)"`
}

type ProductDetails struct {
	Products
	Attributes      map[string]string `gorm:"column:attributes"`
	Tags            []string          `gorm:"column:tags"`
	RelatedChannels []string          `gorm:"column:related_channels"`
	RelatedContents []string          `gorm:"column:related_contents"`
	Reviews         []ProductReviews  `gorm:"column:reviews"`
	ShipData        ShippingData      `gorm:"column:shipping_data"`
}

func (Products) TableName() string {
	return "store_products"
}
