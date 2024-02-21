package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryCustomer struct {
	DB *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *RepositoryCustomer {
	return &RepositoryCustomer{DB: db}
}

func (repository *RepositoryCustomer) ReadAddressByCustomerID(modelAddrs *[]models.CustomerAddresses, customerID uint64) {
	repository.DB.Where("customer_id = ? And active = 1", customerID).Find(modelAddrs)
}
