package zonesvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingZone(db *gorm.DB) *Service {
	return &Service{DB: db}
}
