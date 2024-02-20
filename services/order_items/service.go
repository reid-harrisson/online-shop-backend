package orditmsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceOrderItem(db *gorm.DB) *Service {
	return &Service{DB: db}
}
