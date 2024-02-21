package etsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceEmailTemplate(db *gorm.DB) *Service {
	return &Service{DB: db}
}
