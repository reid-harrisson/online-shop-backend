package models

import (
	"OnlineStoreBackend/pkgs/utils"

	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model

	ProductID  uint64               `gorm:"column:product_id; type:bigint(20) unsigned"`
	CustomerID uint64               `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Comment    string               `gorm:"column:comment; type:text"`
	Rate       float64              `gorm:"column:rate; type:decimal(20,6)"`
	Status     utils.ReviewStatuses `gorm:"column:status; type:tinyint(4)"`
}

func (Reviews) TableName() string {
	return "store_product_reviews"
}
