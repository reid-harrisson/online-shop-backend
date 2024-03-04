package revsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductReview(db *gorm.DB) *Service {
	return &Service{DB: db}
}
