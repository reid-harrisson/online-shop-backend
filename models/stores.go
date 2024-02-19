package models

import "github.com/jinzhu/gorm"

type Stores struct {
	gorm.Model

	CompanyID            uint64 `gorm:"column:company_id; type:bigint(20) unsigned"`
	OwnerID              uint64 `gorm:"column:owner_id; type:bigint(20) unsigned"`
	ContactPhone         string `gorm:"column:contact_phone; type:varchar(25)"`
	ContactEmail         string `gorm:"column:contact_email; type:varchar(100)"`
	ShowStockLevelStatus int8   `gorm:"column:show_stock_level_status; type:tinyint(4)"`
	ShowOutOfStockStatus int8   `gorm:"column:show_out_of_stock_status; type:tinyint(4)"`
	DeliveryPolicy       string `gorm:"column:delivery_policy; type:text"`
	ReturnsPolicy        string `gorm:"column:returns_policy; type:text"`
	Terms                string `gorm:"column:terms; type:text"`
	IsBackOrder          int8   `gorm:"column:is_back_order; type:tinyint(4)"`
	Active               int8   `gorm:"column:active; type:tinyint(4)"`
}

func (Stores) TableName() string {
	return "stores"
}
