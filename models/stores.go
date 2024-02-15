package models

import "github.com/jinzhu/gorm"

type Stores struct {
	gorm.Model

	CompanyID              uint64  `gorm:"column:company_id; type:bigint(20) unsigned"`
	UserID                 uint64  `gorm:"column:user_id; type:bigint(20) unsigned"`
	ContactPhone           string  `gorm:"column:contact_phone; type:varchar(25)"`
	ContactEmail           string  `gorm:"column:contact_email; type:varchar(100)"`
	ShowStockQuantity      int8    `gorm:"column:show_stock_quantity; type:tinyint(4)"`
	ShowOutOfStockProducts int8    `gorm:"column:show_out_of_stock_products; type:tinyint(4)"`
	DeliveryPolicy         string  `gorm:"column:delivery_policy; type:text"`
	ReturnsPolicy          string  `gorm:"column:returns_policy; type:text"`
	Terms                  string  `gorm:"column:terms; type:text"`
	FlatRateShipping       float64 `gorm:"column:flat_rate_shipping; type:decimal(20,6)"`
	BackOrder              int8    `gorm:"column:back_order; type:tinyint(4)"`
	Active                 int8    `gorm:"column:active; type:tinyint(4)"`
}

func (Stores) TableName() string {
	return "stores"
}
