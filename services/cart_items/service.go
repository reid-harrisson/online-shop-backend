package cartsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceCartItem(db *gorm.DB) *Service {
	return &Service{DB: db}
}
