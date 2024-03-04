package prodtagsvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductTag(db *gorm.DB) *Service {
	return &Service{DB: db}
}
