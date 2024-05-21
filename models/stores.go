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
	Active               int8                 `gorm:"column:active; type:tinyint(8)"`
	BackgroundColor1     string               `gorm:"column:background_color_1; type:varchar(25)"`
	BackgroundColor2     string               `gorm:"column:background_color_2; type:varchar(25)"`
	StoreBackground      string               `gorm:"column:store_background; type:varchar(255)"`
	StoreLogo            string               `gorm:"column:store_logo; type:varchar(255)"`
	Description          string               `gorm:"column:description; type:text"`
	HeaderLayoutStyle    int8                 `gorm:"column:header_layout_style; type:tinyint(8)"`
	ShowStoreLogo        int8                 `gorm:"column:show_store_logo; type:tinyint(8)"`
	ShowStoreTitleText   int8                 `gorm:"column:show_store_title_text; type:tinyint(8)"`
	Website              string               `gorm:"column:website; type:varchar(150)"`
	WebsiteButtonColor   string               `gorm:"column:website_button_color; type:varchar(25)"`
}

func (Stores) TableName() string {
	return "stores"
}

func (model *Stores) AfterDelete(db *gorm.DB) (err error) {
	var modelAttrs = []Categories{}
	db.Where("store_id = ?", model.ID).Find(&modelAttrs)
	if len(modelAttrs) > 0 {
		db.Delete(&modelAttrs)
	}

	var modelCombos = []Combos{}
	db.Where("store_id = ?", model.ID).Find(&modelCombos)
	if len(modelCombos) > 0 {
		db.Delete(&modelCombos)
	}

	var modelTemps = []EmailTemplates{}
	db.Where("store_id = ? And link_id = ?", model.ID, model.ID).Find(&modelTemps)
	if len(modelTemps) > 0 {
		db.Delete(&modelTemps)
	}

	var modelItems = []OrderItems{}
	db.Where("store_id = ?", model.ID).Find(&modelItems)
	if len(modelItems) > 0 {
		db.Delete(&modelItems)
	}

	var modelProds = []Products{}
	db.Where("product_id = ?", model.ID).Find(&modelProds)
	if len(modelProds) > 0 {
		db.Delete(&modelProds)
	}

	db.Where("badge_id = ?", model.ID).Delete("invoice_item")

	return
}
