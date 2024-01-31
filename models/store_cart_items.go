package models

import "github.com/jinzhu/gorm"

type CartItems struct {
	gorm.Model

	ProductID  uint64  `gorm:"column:store_product_id; type:bigint(20) unsigned"`
	StoreID    uint64  `gorm:"column:store_id; type:bigint(20) unsigned"`
	CustomerID uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Quantity   float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

type CartItemWithPrice struct {
	CartItems
	UnitPriceSale float64 `gorm:"column:unit_price_sale; type:decimal(20,6)"`
	Price         float64 `gorm:"column:price; type:decimal(20,6)"`
}

type TaxSettings struct {
	TaxRate   float64 `gorm:"column:tax_rate; type:decimal(20,6)"`
	CountryID uint64  `gorm:"column:country_id; type:bigint(20) unsigned"`
}

func (CartItems) TableName() string {
	return "store_cart_items"
}
