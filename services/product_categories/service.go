package prodcatesvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductCategory(db *gorm.DB) *Service {
	return &Service{DB: db}
}
