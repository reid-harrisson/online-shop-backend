package ratesvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductRate(db *gorm.DB) *Service {
	return &Service{DB: db}
}
