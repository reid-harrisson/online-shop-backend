package contsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductContent(db *gorm.DB) *Service {
	return &Service{DB: db}
}
