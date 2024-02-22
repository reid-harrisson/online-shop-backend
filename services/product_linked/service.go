package linkedsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductLinked(db *gorm.DB) *Service {
	return &Service{DB: db}
}
