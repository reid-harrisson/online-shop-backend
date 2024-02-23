package prodvardetsvc

import "github.com/jinzhu/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceProductVariationDetail(db *gorm.DB) *Service {
	return &Service{DB: db}
}
