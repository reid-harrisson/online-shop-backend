package chansvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductChannel(db *gorm.DB) *Service {
	return &Service{DB: db}
}
