package tagsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceTag(db *gorm.DB) *Service {
	return &Service{DB: db}
}
