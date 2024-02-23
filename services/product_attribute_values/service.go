package prodattrvalsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductVariation(db *gorm.DB) *Service {
	return &Service{DB: db}
}
