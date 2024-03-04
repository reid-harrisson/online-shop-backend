package prodattrvalsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductAttributeValue(db *gorm.DB) *Service {
	return &Service{DB: db}
}
