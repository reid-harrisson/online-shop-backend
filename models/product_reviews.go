package models

import "github.com/jinzhu/gorm"

type ProductReviews struct {
	gorm.Model

	ProductID  uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	CustomerID uint64 `gorm:"column:customer_id; type:bigint(20) unsigned"`
	Comment    string `gorm:"column:comment; type:varchar(500)"`
	Status     string `gorm:"column:status; type:varchar(20)"`
}

func (ProductReviews) TableName() string {
	return "store_product_reviews"
}
