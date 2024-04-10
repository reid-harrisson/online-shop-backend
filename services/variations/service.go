package prodvarsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceVariation(db *gorm.DB) *Service {
	return &Service{DB: db}
}
