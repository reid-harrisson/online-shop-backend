package shipoptsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingOption(db *gorm.DB) *Service {
	return &Service{DB: db}
}
