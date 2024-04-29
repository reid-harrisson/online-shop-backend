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
	readVariationsOutputs = []models.Variations{
		{
			ProductID:       1,
			Sku:             "44-125G",
			Price:           96.00,
			StockLevel:      10.00,
			DiscountAmount:  20.00,
			DiscountType:    1,
			ImageUrls:       "",
			Description:     "Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.",
			Title:           "Gochujang - Korean Chilli Pepper Paste - 125G",
			BackOrderStatus: 0,
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

func TestReadVariationsByID(t *testing.T) {
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
	varRepo := repositories.NewRepositoryVariation(db)
	modelVars := models.Variations{}

	// Assertions
	if assert.NoError(t, varRepo.ReadByID(&modelVars, 1)) {
		readVariationsOutputs[0].ID = modelVars.ID
		readVariationsOutputs[0].CreatedAt = modelVars.CreatedAt
		readVariationsOutputs[0].UpdatedAt = modelVars.UpdatedAt

		assert.Equal(t, readVariationsOutputs[0], modelVars)
	}
}

func TestReadVariationsBySku(t *testing.T) {
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
	varRepo := repositories.NewRepositoryVariation(db)
	modelVars := models.Variations{}

	// Assertions
	if assert.NoError(t, varRepo.ReadBySku(&modelVars, "44-125G")) {
		readVariationsOutputs[0].ID = modelVars.ID
		readVariationsOutputs[0].CreatedAt = modelVars.CreatedAt
		readVariationsOutputs[0].UpdatedAt = modelVars.UpdatedAt

		assert.Equal(t, readVariationsOutputs[0], modelVars)
	}
}

func TestReadByAttributeValueIDs(t *testing.T) {
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
	varRepo := repositories.NewRepositoryVariation(db)
	modelVars := models.Variations{}

	// Assertions
	var valueIDs = []uint64{1}
	if assert.NoError(t, varRepo.ReadByAttributeValueIDs(&modelVars, valueIDs, 1)) {
		readVariationsOutputs[0].ID = modelVars.ID
		readVariationsOutputs[0].CreatedAt = modelVars.CreatedAt
		readVariationsOutputs[0].UpdatedAt = modelVars.UpdatedAt

		assert.Equal(t, readVariationsOutputs[0], modelVars)
	}
}
