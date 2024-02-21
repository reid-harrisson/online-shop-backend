package models

import "github.com/jinzhu/gorm"

type CustomerAddresses struct {
	gorm.Model

	ProductID  uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	CustomerID string `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Address    string `gorm:"column:address; type:varchar(100)"`
}

func (CustomerAddresses) TableName() string {
	return "store_customer_addresses"
}
