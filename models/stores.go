package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type Stores struct {
	gorm.Model

	CompanyID            uint64               `gorm:"column:company_id; type:bigint(20) unsigned"`
	OwnerID              uint64               `gorm:"column:owner_id; type:bigint(20) unsigned"`
	Name                 string               `gorm:"column:name; type:varchar(100)"`
	ContactPhone         string               `gorm:"column:contact_phone; type:varchar(25)"`
	ContactEmail         string               `gorm:"column:contact_email; type:varchar(100)"`
	ShowStockLevelStatus utils.SimpleStatuses `gorm:"column:show_stock_level_status; type:tinyint(4)"`
	ShowOutOfStockStatus utils.SimpleStatuses `gorm:"column:show_out_of_stock_status; type:tinyint(4)"`
	DeliveryPolicy       string               `gorm:"column:delivery_policy; type:text"`
	ReturnsPolicy        string               `gorm:"column:returns_policy; type:text"`
	Terms                string               `gorm:"column:terms; type:text"`
}

func (Stores) TableName() string {
	return "stores"
}

func (model *Stores) AfterDelete(db *gorm.DB) (err error) {
	var modelAttrs = []Categories{}
	db.Where("store_id = ?", model.ID).Find(&modelAttrs)
	db.Delete(&modelAttrs)

	var modelCombos = []Combos{}
	db.Where("store_id = ?", model.ID).Find(&modelCombos)
	db.Delete(&modelCombos)

	var modelTemps = []EmailTemplates{}
	db.Where("store_id = ? And link_id = ?", model.ID, model.ID).Find(&modelTemps)
	db.Delete(&modelTemps)

	var modelItems = []OrderItems{}
	db.Where("store_id = ?", model.ID).Find(&modelItems)
	db.Delete(&modelItems)

	var modelProds = []Products{}
	db.Where("product_id = ?", model.ID).Find(&modelProds)
	db.Delete(&modelProds)

	db.Where("badge_id = ?", model.ID).Delete("invoice_item")

	return
}
