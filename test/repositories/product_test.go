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
	readProductsWithDetails = []models.ProductsWithDetail{
		{},
	}
)

func TestReadDetailProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)

	// Setup
	modelProduct := models.ProductsWithDetail{}
	prodRepo := repositories.NewRepositoryProduct(db)

	// Assertions
	if assert.NoError(t, prodRepo.ReadDetail(&modelProduct, 1)) {

		assert.Equal(t, readProductsWithDetails[0], modelProduct)
	}
}
