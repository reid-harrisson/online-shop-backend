package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type Variations struct {
	gorm.Model

	ProductID       uint64               `gorm:"column:product_id; type:bigint(20) unsigned"`
	Sku             string               `gorm:"column:sku; type:varchar(50)"`
	Price           float64              `gorm:"column:price; type:decimal(20,6)"`
	StockLevel      float64              `gorm:"column:stock_level; type:decimal(20,6)"`
	DiscountAmount  float64              `gorm:"column:discount_amount; type:decimal(20,6)"`
	DiscountType    utils.DiscountTypes  `gorm:"column:discount_type; type:tinyint(4)"`
	ImageUrls       string               `gorm:"column:image_urls; type:text"`
	Description     string               `gorm:"column:description; type:text"`
	Title           string               `gorm:"column:title; type:varchar(100)"`
	BackOrderStatus utils.SimpleStatuses `gorm:"column:back_order_status; type:tinyint(4)"`
}

func (Variations) TableName() string {
	return "store_product_variations"
}

type VariationsWithAttributeValue struct {
	Variations
	AttributeValueID uint64 `gomr:"column:attribute_value_id"`
	AttributeName    string `gorm:"column:attribute_name"`
	AttributeValue   string `gorm:"column:attribute_value"`
	Unit             string `gorm:"column:unit"`
}

func (model *Variations) AfterDelete(db *gorm.DB) (err error) {
	var modelCartItems = []CartItems{}
	db.Where("variation_id = ?", model.ID).Find(&modelCartItems)
	if len(modelCartItems) > 0 {
		db.Delete(&modelCartItems)
	}

	var modelComboItems = []ComboItems{}
	db.Where("variation_id = ?", model.ID).Find(&modelComboItems)
	if len(modelComboItems) > 0 {
		db.Delete(&modelComboItems)
	}

	var modelOrderItems = []ComboItems{}
	db.Where("variation_id = ?", model.ID).Find(&modelOrderItems)
	if len(modelOrderItems) > 0 {
		db.Delete(&modelOrderItems)
	}

	var modelDets = []VariationDetails{}
	db.Where("variation_id = ?", model.ID).Find(&modelDets)
	if len(modelDets) > 0 {
		db.Delete(&modelDets)
	}

	var modelShipData = []ShippingData{}
	db.Where("variation_id = ?", model.ID).Find(&modelShipData)
	if len(modelShipData) > 0 {
		db.Delete(&modelShipData)
	}

	return
}
