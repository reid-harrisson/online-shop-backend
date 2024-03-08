package tablesvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingTableRate(db *gorm.DB) *Service {
	return &Service{DB: db}
}
