package shipsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingData(db *gorm.DB) *Service {
	return &Service{DB: db}
}
