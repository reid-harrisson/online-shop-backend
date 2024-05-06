package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodvardetsvc "OnlineStoreBackend/services/variation_details"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	detInputs = []models.VariationDetails{
		{
			VariationID:      1,
			AttributeValueID: 1,
		},
		{
			VariationID:      2,
			AttributeValueID: 2,
		},
	}
	detMatches = []string{
		"1:1",
		"2:2",
	}
	detIndices = map[string]int{
		"1:1": 0,
		"2:2": 1,
	}
)

func TestCreateVariationDetailsWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetAttributeValuesDB(db)
	test_utils.ResetVariationDetailsDB(db)

	// Setup
	var detService = prodvardetsvc.NewServiceVariationDetail(db)
	var modelVars = detInputs

	// Assertions
	if assert.NoError(t, detService.CreateWithCSV(&modelVars, detMatches, detIndices)) {
		assert.Equal(t, modelVars[0].ID, uint(1))
		assert.Equal(t, modelVars[1].ID, uint(2))
	}
}
