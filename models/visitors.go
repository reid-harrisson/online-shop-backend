package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type Visitors struct {
	gorm.Model

	StoreID   uint64          `gorm:"column:store_id; type:bigint(20) unsigned"`
	IpAddress string          `gorm:"column:ip_address; type:varchar(50)"`
	Page      utils.PageTypes `gorm:"column:page; type:tinyint(4)"`
	Bounce    uint64          `gorm:"column:bounce; type:bigint(20)"`
}

func (Visitors) TableName() string {
	return "store_visitors"
}
