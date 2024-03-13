package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type Combos struct {
	gorm.Model

	StoreID        uint64              `gorm:"column:store_id; type:bigint(20)"`
	DiscountAmount float64             `gorm:"column:discount_amount; type:decimal(20,6)"`
	DiscountType   utils.DiscountTypes `gorm:"column:discount_type; type:tinyint(4)"`
	ImageUrls      string              `gorm:"column:image_urls; type:text"`
	Description    string              `gorm:"column:description; type:text"`
	Title          string              `gorm:"column:title; type:varchar(100)"`
}

type ComboItems struct {
	gorm.Model

	ComboID     uint64  `gorm:"column:combo_id; type:bigint(20)"`
	VariationID uint64  `gorm:"column:variation_id; type:bigint(20)"`
	Quantity    float64 `gorm:"column:quantity; type:decimal(20,6)"`
}

func (Combos) TableName() string {
	return "store_combos"
}

func (ComboItems) TableName() string {
	return "store_combo_items"
}
