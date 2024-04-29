package models

import "gorm.io/gorm"

type Addresses struct {
	gorm.Model

	Name         string `gorm:"column:name; type:varchar(255)"`
	AddressLine1 string `gorm:"column:address_line1; type:varchar(100)"`
	AddressLine2 string `gorm:"column:address_line2; type:varchar(100)"`
	SubUrb       string `gorm:"column:suburb; type:varchar(50)"`
	CustomerID   uint64 `gorm:"column:customer_id; type:bigint(20) unsigned"`
	CountryID    uint64 `gorm:"column:country_id; type:bigint(20) unsigned"`
	RegionID     uint64 `gorm:"column:region_id; type:bigint(20) unsigned"`
	CityID       uint64 `gorm:"column:city_id; type:bigint(20) unsigned"`
	PostalCode   string `gorm:"column:postal_code; type:varchar(20)"`
	Active       int8   `gorm:"column:active"`
}

func (Addresses) TableName() string {
	return "store_customer_addresses"
}

func (model *Addresses) AfterDelete(db *gorm.DB) (err error) {
	var modelOrders = []Orders{}
	db.Where("billing_address_id = ? Or shipping_address_id = ?", model.ID, model.ID).Find(&modelOrders)
	db.Delete(&modelOrders)

	return
}
