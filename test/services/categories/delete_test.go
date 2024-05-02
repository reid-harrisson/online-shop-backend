package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	catesvc "OnlineStoreBackend/services/categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteProducts(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetCategoriesDB(db)

	// Setup
	cateService := catesvc.NewServiceCategory(db)

	// Assertions
	assert.NoError(t, cateService.Delete(1))
}
