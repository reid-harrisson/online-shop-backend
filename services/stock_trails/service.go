package stocksvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceStockTrail(db *gorm.DB) *Service {
	return &Service{DB: db}
}
