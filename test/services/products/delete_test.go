package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	linksvc "OnlineStoreBackend/services/links"
	prodsvc "OnlineStoreBackend/services/products"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteProducts(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var prodService = prodsvc.NewServiceProduct(db)

	// Assertions
	assert.NoError(t, prodService.Delete(1))
}

func TestDeleteByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var valService = prodattrvalsvc.NewServiceAttributeValue(db)

	// Assertions
	assert.NoError(t, valService.DeleteByID(1))
}

func TestDeleteProductsAttribute(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var attrService = prodattrsvc.NewServiceAttribute(db)

	// Assertions
	assert.NoError(t, attrService.Delete(1))
}

func TestDeleteProductsShipping(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var shipService = shipsvc.NewServiceShippingData(db)

	// Assertions
	assert.NoError(t, shipService.Delete(1))
}

func TestDeleteProductsLink(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	var linkService = linksvc.NewServiceLink(db)

	// Assertions
	assert.NoError(t, linkService.Delete(1))
}
