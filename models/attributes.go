package models

import "gorm.io/gorm"

type Attributes struct {
	gorm.Model

	ProductID     uint64 `gorm:"column:product_id; type:bigint(20) unsigned"`
	AttributeName string `gorm:"column:attribute_name; type:varchar(50)"`
}

func (Attributes) TableName() string {
	return "store_product_attributes"
}

func (model *Attributes) AfterDelete(db *gorm.DB) (err error) {
	var modelVals = []AttributeValues{}
	db.Where("attribute_id = ?", model.ID).Find(&modelVals)
	if len(modelVals) > 0 {
		db.Delete(&modelVals)
	}

	return
}
