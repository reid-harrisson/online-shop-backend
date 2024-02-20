package ordsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceOrder(db *gorm.DB) *Service {
	return &Service{DB: db}
}
