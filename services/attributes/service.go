package prodattrsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductAttribute(db *gorm.DB) *Service {
	return &Service{DB: db}
}
