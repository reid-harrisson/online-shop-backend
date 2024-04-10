package prodattrsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceAttribute(db *gorm.DB) *Service {
	return &Service{DB: db}
}
