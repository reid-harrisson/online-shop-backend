package addrsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceAddress(db *gorm.DB) *Service {
	return &Service{DB: db}
}
