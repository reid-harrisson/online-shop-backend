package models

import "github.com/jinzhu/gorm"

type Products struct {
	gorm.Model

	CompanyID        uint64  `gorm:"column:company_id; type:bigint(20) unsinged"`
	UserID           uint64  `gorm:"column:user_id; type:bigint(20) unsinged"`
	Name             string  `gorm:"column:name; type:varchar(100)"`
	Brief            string  `gorm:"column:brief; type:varchar(100)"`
	Description      string  `gorm:"column:description; type:varchar(500)"`
	ImageUrls        string  `gorm:"column:image_urls; type:varchar(1000)"`
	SKU              string  `gorm:"column:sku; type:varchar(45)"`
	Tags             string  `gorm:"column:tags; type:varchar(200)"`
	UnitPriceRegular float64 `gorm:"column:unit_price_regular; type:decimal(20,6)"`
	UnitPriceSale    float64 `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	StockQuantity    float64 `gorm:"column:stock_quantity; type:decimal(20,6)"`
	ShippingDataID   uint64  `gorm:"column:shipping_data_id; type:bigint(20) unsinged"`
	LinkedProductIDs string  `gorm:"column:linked_product_ids; type:varchar(100)"`
	Attributes       string  `gorm:"column:attributes; type:varchar(500)"`
	Active           int8    `gorm:"column:active; type:tinyint(4)"`
	Reviews          string  `gorm:"column:reviews; type:text"`
}

func (Products) TableName() string {
	return "store_products"
}

type ProductDetails struct {
	Products
	ImgUrls      []string
	Attribs      map[string]string
	ShippingInfo ShippingData
	ChannelIDs   string
	ContentIDs   string
}
