package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	addrOutputs = []models.Addresses{
		{
			Name:         "",
			AddressLine1: "Andorra, Andorra",
			AddressLine2: "",
			SubUrb:       "Andorra la Vella",
			CountryID:    1,
			RegionID:     1,
			CityID:       1,
			PostalCode:   "11-111",
			CustomerID:   1,
			Active:       1,
		},
	}
)

func TestReadAddressByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAddressesDB(db)
	test_utils.ResetCitiesDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	addrRepo := repositories.NewRepositoryAddresses(db)
	var modelAddress = models.Addresses{}

	// Assertions
	if assert.NoError(t, addrRepo.ReadByID(&modelAddress, 1, 1)) {
		addrOutputs[0].Model.ID = modelAddress.Model.ID
		addrOutputs[0].CreatedAt = modelAddress.CreatedAt
		addrOutputs[0].UpdatedAt = modelAddress.UpdatedAt

		assert.Equal(t, addrOutputs[0], modelAddress)
	}
}

func TestReadAddressesByCustomerID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetAddressesDB(db)
	test_utils.ResetCitiesDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	addrRepo := repositories.NewRepositoryAddresses(db)
	var modelAddrs = []models.Addresses{}

	// Assertions
	if assert.NoError(t, addrRepo.ReadAddressesByCustomerID(&modelAddrs, 1)) {
		addrOutputs[0].Model.ID = modelAddrs[0].Model.ID
		addrOutputs[0].CreatedAt = modelAddrs[0].CreatedAt
		addrOutputs[0].UpdatedAt = modelAddrs[0].UpdatedAt

		assert.Equal(t, addrOutputs[0], modelAddrs[0])
	}
}
