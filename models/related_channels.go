package models

import "gorm.io/gorm"

type ProductChannels struct {
	gorm.Model

	ProductID uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	ChannelID uint64 `gorm:"column:channel_id; type:bigint(20) unsigned"`
}

type ProductChannelsWithName struct {
	ProductChannels
	ChannelName string `gorm:"column:channel_name"`
}

func (ProductChannels) TableName() string {
	return "store_product_related_channels"
}
