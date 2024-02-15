package attrsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceAttribute(db *gorm.DB) *Service {
	return &Service{DB: db}
}
