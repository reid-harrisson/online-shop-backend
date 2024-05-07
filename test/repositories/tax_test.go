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
	readTaxes = []models.Taxes{
		{
			TaxRate:    4.5,
			CountryID:  1,
			CustomerID: 1,
		},
	}
)

func TestReadCurrencyTax(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)

	// Setup
	taxRepo := repositories.NewRepositoryTax(db)
	currencyCode := ""
	exchangeRate := float64(0)

	// Assertions
	assert.NoError(t, taxRepo.ReadCurrency(&currencyCode, &exchangeRate, 34))
}

func TestReadByCountryIDTax(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCountriesDB(db)

	// Setup
	taxRepo := repositories.NewRepositoryTax(db)
	modelTax := models.Taxes{}

	// Assertions
	if assert.NoError(t, taxRepo.ReadByCountryID(&modelTax, 1)) {
		assert.Equal(t, readTaxes[0].TaxRate, modelTax.TaxRate)
		assert.Equal(t, readTaxes[0].CountryID, modelTax.CountryID)
	}
}

func TestReadByCustomerIDTax(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCountriesDB(db)

	// Setup
	taxRepo := repositories.NewRepositoryTax(db)
	modelTax := models.Taxes{}

	// Assertions
	if assert.NoError(t, taxRepo.ReadByCustomerID(&modelTax, 1)) {
		assert.Equal(t, readTaxes[0], modelTax)
	}
}
