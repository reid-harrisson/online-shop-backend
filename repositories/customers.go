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

func (repository *RepositoryCustomer) ReadAddressesByCustomerID(modelAddrs *[]models.CustomerAddresses, customerID uint64) {
	repository.DB.Where("customer_id = ? And active = 1", customerID).Find(modelAddrs)
}

func (repository *RepositoryCustomer) ReadAddressByCustomerID(modelAddr *models.CustomerAddresses, customerID uint64) {
	repository.DB.Where("customer_id = ? And active = 1", customerID).First(modelAddr)
}

func (repository *RepositoryCustomer) ReadAddressByID(modelAddr *models.CustomerAddresses, addressID uint64) {
	repository.DB.First(modelAddr, addressID)
}
