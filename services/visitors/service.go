package vistsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceVisitor(db *gorm.DB) *Service {
	return &Service{DB: db}
}
