package flatsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingFlatRate(db *gorm.DB) *Service {
	return &Service{DB: db}
}
