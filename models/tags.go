package models

import "github.com/jinzhu/gorm"

type BaseTags struct {
	gorm.Model

	StoreID uint64 `gorm:"column:store_id; type:bigint(20) unsigned"`
	Name    string `gorm:"column:name; type:varchar(50)"`
}

type ProductTags struct {
	gorm.Model

	ProductID uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	TagID     uint64 `gorm:"column:tag_id; type:bigint(20) unsigned"`
}

func (BaseTags) TableName() string {
	return "store_tags"
}

func (ProductTags) TableName() string {
	return "store_product_tags"
}

type ProductTagsWithName struct {
	ProductTags
	TagName string `gorm:"column:tag_name"`
}
