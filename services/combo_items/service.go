package coitmsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceComboItem(db *gorm.DB) *Service {
	return &Service{DB: db}
}
