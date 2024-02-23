package addrsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceCustomerAddress(db *gorm.DB) *Service {
	return &Service{DB: db}
}
