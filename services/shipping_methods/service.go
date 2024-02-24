package shipmthsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceShippingMethod(db *gorm.DB) *Service {
	return &Service{DB: db}
}
