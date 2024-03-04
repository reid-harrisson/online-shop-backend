package storesvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceStore(db *gorm.DB) *Service {
	return &Service{DB: db}
}
