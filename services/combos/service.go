package combsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceCombo(db *gorm.DB) *Service {
	return &Service{DB: db}
}
