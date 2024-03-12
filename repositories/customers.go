package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryCustomer struct {
	DB *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *RepositoryCustomer {
	return &RepositoryCustomer{DB: db}
}

func (repository *RepositoryCustomer) ReadAddressesByCustomerID(modelAddrs *[]models.Addresses, customerID uint64) {
	repository.DB.Where("customer_id = ? And active = 1", customerID).Find(modelAddrs)
}

func (repository *RepositoryCustomer) ReadAddressByCustomerID(modelAddr *models.Addresses, customerID uint64) error {
	return repository.DB.Where("customer_id = ? And active = 1", customerID).First(modelAddr).Error
}

func (repository *RepositoryCustomer) ReadAddressByID(modelAddr *models.Addresses, addressID uint64) {
	repository.DB.First(modelAddr, addressID)
}
