package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

func TestUpdateByIDAttributeValues(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	valService := prodattrvalsvc.NewServiceAttributeValue(db)

	// Assertions
	assert.NoError(t, valService.UpdateByID(1, "10"))
}
