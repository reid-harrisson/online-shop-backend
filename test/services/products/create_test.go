package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
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
	reqProductWithDetail = []requests.RequestProductWithDetail{
		{
			RequestProduct: requests.RequestProduct{
				StoreID:          1,
				Title:            "1",
				ShortDescription: "1",
				LongDescription:  "1",
				ImageUrls:        []string{"1", "1"},
			},

			RelatedChannels: []uint64{1, 2},
			RelatedContents: []uint64{1, 2},
			Categories:      []uint64{1, 2},
			Tags:            []string{"kefir"},
			Attributes: map[string][]string{
				"additionalProp1": {"string"},
				"additionalProp2": {"string"},
				"additionalProp3": {"string"},
			},
			UpSell:    []uint64{1, 2},
			CrossSell: []uint64{1, 2},
		},
	}
)

func TestCreateProducts(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)
	var modelProducts = productInputs

	// Assertions
	if assert.NoError(t, prodService.Create(&modelProducts[0], &reqProductWithDetail[0])) {
		assert.Equal(t, modelProducts, productInputs)
	}
}

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
