package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodsvc "OnlineStoreBackend/services/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	productInputs = []models.Products{
		{
			StoreID:           1,
			Title:             "Natural Milk Kefir",
			ShortDescription:  "Short description for natural milk kefir",
			LongDescription:   "Long description for natural milk kefir",
			ImageUrls:         "https://www.chegourmet.co.za/wp-content/uploads/2022/09/1.png",
			MinimumStockLevel: 0,
			Status:            0,
			Sku:               "57",
			Type:              0,
			ShippingClass:     "Courier Refrigerated",
		},
	}
)

func TestCreateProductsWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)
	var modelProducts = productInputs

	// Assertions
	if assert.NoError(t, prodService.CreateWithCSV(&modelProducts, []string{"57"}, map[string]int{"57": 0})) {
		assert.Equal(t, modelProducts, productInputs)
	}
}
