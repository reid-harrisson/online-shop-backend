package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

func TestReadCurrencyTax(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetShippingData(db)

	// Setup
	taxRepo := repositories.NewRepositoryTax(db)
	currencyCode := ""
	exchangeRate := float64(0)

	// Assertions
	assert.NoError(t, taxRepo.ReadCurrency(&currencyCode, &exchangeRate, 34))
}
