package storesvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceStore(db *gorm.DB) *Service {
	return &Service{DB: db}
}
