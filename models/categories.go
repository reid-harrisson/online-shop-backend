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

type ProductCategories struct {
	gorm.Model

	ProductID  uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	CategoryID uint64 `gorm:"column:category_id; type:bigint(20) unsigned"`
}

func (BaseCategories) TableName() string {
	return "store_categories"
}

func (ProductCategories) TableName() string {
	return "store_product_categories"
}

type ProductCategoriesWithName struct {
	ProductCategories
	CategoryName string `gorm:"column:category_name"`
}
