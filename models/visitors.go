package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"github.com/jinzhu/gorm"
)

type Visitors struct {
	gorm.Model

	StoreID     uint64          `gorm:"column:store_id; type:bigint(20) unsigned"`
	ProductID   uint64          `gorm:"column:product_id; type:bigint(20) unsigned"`
	IpAddress   string          `gorm:"column:ip_address; type:varchar(50)"`
	Page        utils.PageTypes `gorm:"column:page; type:tinyint(4)"`
	Bounce      uint64          `gorm:"column:bounce; type:bigint(20)"`
	LoadingTime float64         `gorm:"column:loading_time; type:decimal(20,6)"`
	ErrorCode   int             `gorm:"column:error_code; type:int"`
}

func (Visitors) TableName() string {
	return "store_visitors"
}
