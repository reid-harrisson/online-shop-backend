package models

import "github.com/jinzhu/gorm"

type ProductCustomerRates struct {
	gorm.Model

	ProductID  uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	CustomerID uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Rate       float64 `gorm:"column:rate; type:decimal(20,6)"`
}

type ProductRates struct {
	ProductID uint64  `gorm:"column:product_id; type:bigint(20) unsigned"`
	Rate      float64 `gorm:"column:rate; type:decimal(20,6)"`
	Customers uint64  `gorm:"column:customers; type:bigint(20) unsigned"`
}

func (ProductCustomerRates) TableName() string {
	return "store_product_rates"
}
