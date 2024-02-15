package catesvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceCategory(db *gorm.DB) *Service {
	return &Service{DB: db}
}
