package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type StockTracks struct {
	gorm.Model

	ProductID   uint64            `gorm:"column:product_id; type:bigint(20) unsigned"`
	VariationID uint64            `gorm:"column:variation_id; type:bigint(20) unsigned"`
	Change      float64           `gorm:"column:change; type:decimal(20,6)"`
	Event       utils.TrackEvents `gorm:"column:event; type:tinyint(4)"`
}

func (StockTracks) TableName() string {
	return "store_stock_tracks"
}
