package models

import "gorm.io/gorm"

type AttributeValues struct {
	gorm.Model

	AttributeID    uint64 `gorm:"column:attribute_id; type:bigint(20) unsigned"`
	AttributeValue string `gorm:"column:attribute_value; type:varchar(50)"`
}

func (AttributeValues) TableName() string {
	return "store_product_attribute_values"
}

type AttributeValuesWithDetail struct {
	AttributeValues
	AttributeName string `gorm:"column:attribute_name"`
}

func (model *AttributeValues) AfterDelete(db *gorm.DB) (err error) {
	var modelDets = []VariationDetails{}
	db.Where("attribute_value_id = ?", model.ID).Find(&modelDets)
	db.Delete(&modelDets)

	return
}
