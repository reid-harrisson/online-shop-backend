package models

import "gorm.io/gorm"

type Tags struct {
	gorm.Model

	StoreID uint64 `gorm:"column:store_id; type:bigint(20) unsigned"`
	Name    string `gorm:"column:name; type:varchar(50)"`
}

type ProductTags struct {
	gorm.Model

	ProductID uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	TagID     uint64 `gorm:"column:tag_id; type:bigint(20) unsigned"`
}

func (Tags) TableName() string {
	return "store_tags"
}

func (ProductTags) TableName() string {
	return "store_product_tags"
}

type ProductTagsWithName struct {
	ProductTags
	TagName string `gorm:"column:tag_name"`
}

func (model *Tags) AfterDelete(db *gorm.DB) (err error) {
	var modelTags = []ProductTags{}
	db.Where("tag_id = ?", model.ID).Find(&modelTags)
	if len(modelTags) > 0 {
		db.Delete(&modelTags)
	}

	return
}
