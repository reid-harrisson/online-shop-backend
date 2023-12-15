package models

import "github.com/jinzhu/gorm"

type ProductChannels struct {
	gorm.Model

	ProductID uint64 `gorm:"column:store_product_id; type:bigint(20) unsinged"`
	ChannelID uint64 `gorm:"column:channel_id; type:bigint(20) unsinged"`
}

func (ProductChannels) TableName() string {
	return "store_product_related_channels"
}
