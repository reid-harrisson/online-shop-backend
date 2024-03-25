package stocksvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceStockTrack(db *gorm.DB) *Service {
	return &Service{DB: db}
}
