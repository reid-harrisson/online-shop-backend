package prodattrsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductAttribute(db *gorm.DB) *Service {
	return &Service{DB: db}
}
