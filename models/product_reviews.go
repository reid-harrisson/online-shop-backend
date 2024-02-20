package models

import "github.com/jinzhu/gorm"

type ReviewStatuses int8

const (
	StatusReviewPending ReviewStatuses = iota + 0
	StatusReviewApproved
	StatusReviewBlocked
)

type ProductReviews struct {
	gorm.Model

	ProductID  uint64         `gorm:"column:product_id; type:bigint(20) unsigned"`
	CustomerID uint64         `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Comment    string         `gorm:"column:comment; type:text"`
	Rate       float64        `gorm:"column:rate; type:decimal(20,6)"`
	Status     ReviewStatuses `gorm:"column:status; type:tinyint(4)"`
}

func (ProductReviews) TableName() string {
	return "store_product_reviews"
}
