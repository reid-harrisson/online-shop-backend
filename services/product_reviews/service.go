package revsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductReview(db *gorm.DB) *Service {
	return &Service{DB: db}
}
