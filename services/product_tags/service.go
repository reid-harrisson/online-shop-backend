package prodtagsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductTag(db *gorm.DB) *Service {
	return &Service{DB: db}
}
