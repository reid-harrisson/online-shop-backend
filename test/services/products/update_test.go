package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	prodsvc "OnlineStoreBackend/services/products"
	"testing"

	"github.com/stretchr/testify/assert"
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
