package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	storesvc "OnlineStoreBackend/services/stores"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteStore(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetUsersDB(db)
	test_utils.ResetCompaniesDB(db)
	test_utils.ResetStoresDB(db)

	// Setup
	var storeService = storesvc.NewServiceStore(db)

	// Assertions
	assert.NoError(t, storeService.Delete(1))
}
