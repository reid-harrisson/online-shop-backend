package models

import "gorm.io/gorm"

type StoreCategories struct {
	gorm.Model

	StoreID  uint64 `gorm:"column:store_id; type:bigint(20) unsigned"`
	ParentID uint64 `gorm:"column:parent_id; type:bigint(20) unsigned"`
	Name     string `gorm:"column:name; type:varchar(45)"`
}

type StoreCategoriesWithChildren struct {
	StoreCategories
	ChildrenIDs []uint64 `gorm:"column:children_ids"`
}

type ProductCategories struct {
	gorm.Model

	ProductID  uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	CategoryID uint64 `gorm:"column:category_id; type:bigint(20) unsigned"`
}

func (StoreCategories) TableName() string {
	return "store_categories"
}

func (ProductCategories) TableName() string {
	return "store_product_categories"
}

type ProductCategoriesWithName struct {
	ProductCategories
	CategoryName string `gorm:"column:category_name"`
}
