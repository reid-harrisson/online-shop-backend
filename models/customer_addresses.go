package models

import "github.com/jinzhu/gorm"

type CustomerAddress struct {
	gorm.Model

	CusomterID   uint64 `gorm:"column:customer_id; type:bigint(20) unsigned"`
	CountryID    uint64 `gorm:"column:contry_id; type:bigint(20) unsigned"`
	RegionID     uint64 `gorm:"column:region_id; type:bigint(20) unsigned"`
	CityID       uint64 `gorm:"column:city_id; type:bigint(20) unsigned"`
	PostalCode   string `gorm:"column:postal_code; type:varchar(15)"`
	AddressLine1 string `gorm:"column:address_line1; type:varchar(100)"`
	AddressLine2 string `gorm:"column:address_line2; type:varchar(100)"`
	Suburb       string `gorm:"column:suburb; type:varchar(50)"`
	Active       int8   `gorm:"column; type:tinyint(4) unsigned"`
}
