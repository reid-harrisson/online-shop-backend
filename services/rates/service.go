package ratesvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductRate(db *gorm.DB) *Service {
	return &Service{DB: db}
}
