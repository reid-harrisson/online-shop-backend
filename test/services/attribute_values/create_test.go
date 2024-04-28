package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodattrvalsvc "OnlineStoreBackend/services/attribute_values"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	valInputs = []models.AttributeValues{
		{
			AttributeID:    1,
			AttributeValue: "100g",
		},
	}
)

func TestCreateAttributeValues(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)

	// Setup

	var valService = prodattrvalsvc.NewServiceAttributeValue(db)

	// Assertions
	assert.NoError(t, valService.Create(1, "100g"))
}

func TestCreateAttributeValuesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)

	// Setup

	var valService = prodattrvalsvc.NewServiceAttributeValue(db)
	var modelVals = valInputs

	// Assertions
	if assert.NoError(t, valService.CreateWithCSV(&modelVals, []string{"1:100g"}, map[string]int{"1:100g": 0})) {
		assert.Equal(t, modelVals, valInputs)
	}
}
