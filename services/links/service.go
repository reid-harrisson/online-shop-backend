package linksvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceLink(db *gorm.DB) *Service {
	return &Service{DB: db}
}
