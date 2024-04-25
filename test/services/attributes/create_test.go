package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodattrsvc "OnlineStoreBackend/services/attributes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	attrInputs = []models.Attributes{
		{
			ProductID:     1,
			AttributeName: "color",
		},
	}
)

func TestCreateAttributesWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributesDB(db)

	// Setup
	var attrService = prodattrsvc.NewServiceAttribute(db)
	var modelAttrs = attrInputs

	// Assertions
	if assert.NoError(t, attrService.CreateWithCSV(&modelAttrs, []string{"1:color"}, map[string]int{"1:color": 0})) {
		assert.Equal(t, modelAttrs, attrInputs)
	}
}
