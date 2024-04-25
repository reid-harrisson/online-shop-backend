package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	addrsvc "OnlineStoreBackend/services/addresses"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	addrInput = requests.RequestAddress{
		Name:         "Andorra",
		AddressLine1: "Andorra, Andorra",
		AddressLine2: "",
		SubUrb:       "Andorra la Vella",
		CountryID:    1,
		RegionID:     1,
		CityID:       1,
		PostalCode:   "11-111",
	}
)

func TestCreateAttributeValuesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetUsersDB(db)
	test_utils.ResetAddressesDB(db)

	// Setup

	var addrService = addrsvc.NewServiceAddress(db)
	var modelAddr = models.Addresses{}

	// Assertions
	if assert.NoError(t, addrService.Create(&modelAddr, &addrInput, 1)) {
		assert.Equal(t, uint(3), modelAddr.ID)
		assert.Equal(t, addrInput.Name, modelAddr.Name)
		assert.Equal(t, addrInput.AddressLine1, modelAddr.AddressLine1)
		assert.Equal(t, addrInput.AddressLine2, modelAddr.AddressLine2)
		assert.Equal(t, addrInput.SubUrb, modelAddr.SubUrb)
		assert.Equal(t, addrInput.CountryID, modelAddr.CountryID)
		assert.Equal(t, addrInput.RegionID, modelAddr.RegionID)
		assert.Equal(t, addrInput.CityID, modelAddr.CityID)
		assert.Equal(t, addrInput.PostalCode, modelAddr.PostalCode)
	}
}
