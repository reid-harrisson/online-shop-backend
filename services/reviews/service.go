package revsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceReview(db *gorm.DB) *Service {
	return &Service{DB: db}
}
