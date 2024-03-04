package orditmsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceOrderItem(db *gorm.DB) *Service {
	return &Service{DB: db}
}
