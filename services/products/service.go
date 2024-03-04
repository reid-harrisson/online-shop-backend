package prodsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProduct(db *gorm.DB) *Service {
	return &Service{DB: db}
}
