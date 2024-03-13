package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryAddresses struct {
	DB *gorm.DB
}

func NewRepositoryAddresses(db *gorm.DB) *RepositoryAddresses {
	return &RepositoryAddresses{DB: db}
}

func (repository *RepositoryAddresses) ReadAddressesByCustomerID(modelAddrs *[]models.Addresses, customerID uint64) {
	repository.DB.Where("customer_id = ? And active = 1", customerID).Find(modelAddrs)
}

func (repository *RepositoryAddresses) ReadAddressByCustomerID(modelAddr *models.Addresses, customerID uint64) error {
	return repository.DB.Where("customer_id = ? And active = 1", customerID).First(modelAddr).Error
}

func (repository *RepositoryAddresses) ReadAddressByID(modelAddr *models.Addresses, addressID uint64) {
	repository.DB.First(modelAddr, addressID)
}

func (repository *RepositoryAddresses) ReadByID(modelAddr *models.Addresses, addressID uint64, customerID uint64) {
	repository.DB.Where("id = ? And customer_id = ?", addressID, customerID).First(modelAddr)
}
