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

func TestCreateAttributeValuesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetUsersDB(db)
	test_utils.ResetAddressesDB(db)

	// Setup

	var addrRepo = repositories.NewRepositoryAddresses(db)
	var modelAddrs = []models.Addresses{}

	// Assertions
	if assert.NoError(t, addrRepo.ReadAddressesByCustomerID(&modelAddrs, 1)) {
		if assert.Equal(t, len(addrOutputs), len(modelAddrs)) {
			for index := range modelAddrs {
				addrOutputs[index].CreatedAt = modelAddrs[index].CreatedAt
				addrOutputs[index].UpdatedAt = modelAddrs[index].UpdatedAt
				addrOutputs[index].ID = modelAddrs[index].ID
				assert.Equal(t, addrOutputs[index], modelAddrs[index])
			}
		}
	}
}
