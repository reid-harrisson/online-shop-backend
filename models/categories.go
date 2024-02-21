package models

import "github.com/jinzhu/gorm"

type BaseCategories struct {
	gorm.Model

	StoreID  uint64 `gorm:"column:store_id; type:bigint(20) unsigned"`
	Name     string `gorm:"column:name; type:varchar(50)"`
	ParentID uint64 `gorm:"column:parent_id; type:bigint(20) unsigned"`
}

type StoreCategoriesWithChildren struct {
	BaseCategories
	ChildrenIDs []uint64 `gorm:"column:children_ids"`
}

func (BaseCategories) TableName() string {
	return "store_categories"
}
