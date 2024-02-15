package models

import "github.com/jinzhu/gorm"

type ProductContents struct {
	gorm.Model

	ProductID uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	ContentID uint64 `gorm:"column:content_id; type:bigint(20) unsigned"`
}

type BaseContents struct {
	ID uint64 `gorm:"column:id"`
}

type ProductContentsWithTitle struct {
	ProductContents
	ContentTitle string `gorm:"column:content_title"`
}

func (ProductContents) TableName() string {
	return "store_product_related_contents"
}
