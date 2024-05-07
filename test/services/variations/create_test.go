package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/requests"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	varInputs = []models.Variations{
		{
			ProductID:       1,
			Sku:             "44-125G",
			Price:           96,
			StockLevel:      10,
			DiscountAmount:  20,
			DiscountType:    utils.FixedAmountOff,
			ImageUrls:       "",
			Description:     "Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.",
			Title:           "Gochujang - Korean Chilli Pepper Paste - 125G",
			BackOrderStatus: utils.Disabled,
		},
		{
			ProductID:       2,
			Sku:             "13-200ML",
			Price:           45,
			StockLevel:      10,
			DiscountAmount:  0,
			DiscountType:    utils.PercentageOff,
			ImageUrls:       "",
			Description:     "NuMe Kombucha - 350ml, Buchu, Hibiscus & Hawthorne",
			Title:           "Kimchi Probiotic Tonic - 200ML - 200ML",
			BackOrderStatus: utils.Disabled,
		},
	}
	varMatches = []string{
		"1:44-125G",
		"2:13-200ML",
	}
	varIndices = map[string]int{
		"1:44-125G":  0,
		"2:13-200ML": 1,
	}
	varRequest = requests.RequestVariation{
		Price:             96,
		StockLevel:        10,
		DiscountAmount:    20,
		DiscountType:      utils.FixedAmountOff,
		ImageUrls:         []string{},
		Description:       "Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.",
		BackOrderAllowed:  1,
		AttributeValueIDs: []uint64{1, 2},
	}
	varOutput = models.Variations{
		ProductID:       1,
		Sku:             "GOCHUJANGKOREANCHILLIPEPPERPASTE125G200ML",
		Price:           96,
		StockLevel:      10,
		DiscountAmount:  20,
		DiscountType:    utils.FixedAmountOff,
		ImageUrls:       "[]",
		Description:     "Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.",
		Title:           "Gochujang - Korean Chilli Pepper Paste - 125g, 200ml",
		BackOrderStatus: utils.Enabled,
	}
)

func TestCreateVariation(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVariationDetailsDB(db)
	test_utils.ResetAttributeValuesDB(db)
	test_utils.ResetAttributesDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var varService = prodvarsvc.NewServiceVariation(db)
	var modelVar = models.Variations{}

	// Assertions
	if assert.NoError(t, varService.Create(&modelVar, &varRequest, 1)) {
		varOutput.CreatedAt = modelVar.CreatedAt
		varOutput.UpdatedAt = modelVar.UpdatedAt
		varOutput.ID = 3
		assert.Equal(t, varOutput, modelVar)
	}
}

func TestCreateVariationsWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetAttributeValuesDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var varService = prodvarsvc.NewServiceVariation(db)
	var modelVars = varInputs

	// Assertions
	if assert.NoError(t, varService.CreateWithCSV(&modelVars, varMatches, varIndices)) {
		assert.Equal(t, modelVars[0].ID, uint(1))
		assert.Equal(t, modelVars[1].ID, uint(2))
	}
}
