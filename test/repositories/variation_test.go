package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readVarsWithAttrValue = []models.VariationsWithAttributeValue{
		{
			Variations:       readVars[0],
			AttributeValueID: 1,
			AttributeName:    "weight",
			AttributeValue:   "125g",
			Unit:             "",
		},
	}
)

func TestReadVariationsByProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetAttributeValuesDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetVariationDetailsDB(db)
	test_utils.ResetAttributesDB(db)

	// Setup
	modelVars := []models.VariationsWithAttributeValue{}
	varRepo := repositories.NewRepositoryVariation(db)

	// Assertions
	if assert.NoError(t, varRepo.ReadByProduct(&modelVars, 1)) {
		if assert.Equal(t, len(readVarsWithAttrValue), len(modelVars)) {
			for index := range modelVars {
				readVarsWithAttrValue[index].CreatedAt = modelVars[index].CreatedAt
				readVarsWithAttrValue[index].UpdatedAt = modelVars[index].UpdatedAt
				assert.Equal(t, readVarsWithAttrValue[index], modelVars[index])
			}
		}
	}
}
