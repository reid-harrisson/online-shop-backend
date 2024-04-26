package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readAddresses = []models.CustomerOrdersWithAddress{
		{
			BillingAddress: models.Addresses{
				AddressLine1: "",
				AddressLine2: "",
				SubUrb:       "",
				CustomerID:   1,
				CountryID:    1,
				RegionID:     1,
				CityID:       1,
				PostalCode:   "",
				Active:       1,
			},
			ShippingAddress: models.Addresses{
				AddressLine1: "",
				AddressLine2: "",
				SubUrb:       "",
				CustomerID:   1,
				CountryID:    1,
				RegionID:     1,
				CityID:       1,
				PostalCode:   "",
				Active:       1,
			},
			Items: []models.CustomerOrderItems{},
		},
	}
)

func TestReadAddressByIDBillingAddress(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreCustomerAddressesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	modelOrder := models.CustomerOrdersWithAddress{}
	addrRepo := repositories.NewRepositoryAddresses(db)

	// Assertions
	if assert.NoError(t, addrRepo.ReadAddressByID(&modelOrder.BillingAddress, 1)) {
		readAddresses[0].BillingAddress.Model.ID = modelOrder.BillingAddress.Model.ID
		readAddresses[0].BillingAddress.CreatedAt = modelOrder.BillingAddress.CreatedAt
		readAddresses[0].BillingAddress.UpdatedAt = modelOrder.BillingAddress.UpdatedAt

		assert.Equal(t, readAddresses[0].BillingAddress, modelOrder.BillingAddress)
	}
}

func TestReadAddressByIDShippingAddress(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreCustomerAddressesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	modelOrder := models.CustomerOrdersWithAddress{}
	addrRepo := repositories.NewRepositoryAddresses(db)

	// Assertions
	if assert.NoError(t, addrRepo.ReadAddressByID(&modelOrder.ShippingAddress, 1)) {
		readAddresses[0].ShippingAddress.Model.ID = modelOrder.ShippingAddress.Model.ID
		readAddresses[0].ShippingAddress.CreatedAt = modelOrder.ShippingAddress.CreatedAt
		readAddresses[0].ShippingAddress.UpdatedAt = modelOrder.ShippingAddress.UpdatedAt

		assert.Equal(t, readAddresses[0].ShippingAddress, modelOrder.ShippingAddress)
	}
}
