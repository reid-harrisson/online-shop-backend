package cousvc

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

func NewServiceCoupon(db *gorm.DB) *Service {
	return &Service{DB: db}
}
