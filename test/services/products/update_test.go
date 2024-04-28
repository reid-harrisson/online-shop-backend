package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	prodsvc "OnlineStoreBackend/services/products"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	reqProduct = requests.RequestProduct{
		StoreID:          1,
		Title:            "1",
		ShortDescription: "1",
		LongDescirpiton:  "1",
		ImageUrls:        []string{"1", "2"},
	}
)

func TestUpdateStatus(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)

	// Assertions
	assert.NoError(t, prodService.UpdateStatus(1, utils.Pending))
}

func TestUpdate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)
	modelProduct := models.Products{}

	// Assertions
	assert.NoError(t, prodService.Update(&modelProduct, &reqProduct))
}
