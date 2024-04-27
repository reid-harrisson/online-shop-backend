package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	comboItemsOutputs = []models.ComboItems{
		{
			ComboID:     1,
			VariationID: 1,
			Quantity:    1,
		},
		{
			ComboID:     2,
			VariationID: 2,
			Quantity:    2,
		},
	}
	comboOutputs = []models.Combos{
		{
			StoreID:        1,
			DiscountAmount: 10.000,
			DiscountType:   1,
			ImageUrls:      "https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg",
			Description:    "Combo of Kefir",
			Title:          "Kefir Combo",
			Status:         2,
		},
	}
)

func TestComboReadByStoreID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCombosDB(db)
	test_utils.ResetComboItemsDB(db)
	test_utils.ResetStoresDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var comboRepo = repositories.NewRepositoryCombo(db)

	var modelCombos = []models.Combos{}
	var modelComboItems = []models.ComboItems{}

	// Assertions
	if assert.NoError(t, comboRepo.ReadByStoreID(&modelCombos, &modelComboItems, 1)) {
		assert.Equal(t, comboItemsOutputs[0].ComboID, modelComboItems[0].ComboID)
		assert.Equal(t, comboItemsOutputs[0].VariationID, modelComboItems[0].VariationID)
		assert.Equal(t, comboItemsOutputs[0].ComboID, modelComboItems[0].ComboID)
	}
}

func TestComboReadApproved(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCombosDB(db)
	test_utils.ResetComboItemsDB(db)
	test_utils.ResetStoresDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var comboRepo = repositories.NewRepositoryCombo(db)

	var modelCombos = []models.Combos{}
	var modelComboItems = []models.ComboItems{}

	// Assertions
	if assert.NoError(t, comboRepo.ReadApproved(&modelCombos, &modelComboItems, 1)) {
		assert.Equal(t, comboItemsOutputs[0].ComboID, modelComboItems[0].ComboID)
		assert.Equal(t, comboItemsOutputs[0].VariationID, modelComboItems[0].VariationID)
		assert.Equal(t, comboItemsOutputs[0].ComboID, modelComboItems[0].ComboID)
	}
}

func TestComboReadPublished(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCombosDB(db)
	test_utils.ResetComboItemsDB(db)
	test_utils.ResetStoresDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var comboRepo = repositories.NewRepositoryCombo(db)

	var modelCombos = []models.Combos{}
	var modelComboItems = []models.ComboItems{}

	// Assertions
	if assert.NoError(t, comboRepo.ReadPublished(&modelCombos, &modelComboItems, 2)) {
		assert.Equal(t, comboItemsOutputs[1].ComboID, modelComboItems[0].ComboID)
		assert.Equal(t, comboItemsOutputs[1].VariationID, modelComboItems[0].VariationID)
		assert.Equal(t, comboItemsOutputs[1].ComboID, modelComboItems[0].ComboID)
	}
}

func TestComboReadByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCombosDB(db)
	test_utils.ResetComboItemsDB(db)
	test_utils.ResetStoresDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	var comboRepo = repositories.NewRepositoryCombo(db)

	var modelCombos = models.Combos{}

	// Assertions
	if assert.NoError(t, comboRepo.ReadByID(&modelCombos, 1)) {
		assert.Equal(t, comboOutputs[0].StoreID, modelCombos.StoreID)
		assert.Equal(t, comboOutputs[0].DiscountAmount, modelCombos.DiscountAmount)
		assert.Equal(t, comboOutputs[0].DiscountType, modelCombos.DiscountType)
		assert.Equal(t, comboOutputs[0].ImageUrls, modelCombos.ImageUrls)
		assert.Equal(t, comboOutputs[0].Description, modelCombos.Description)
		assert.Equal(t, comboOutputs[0].Title, modelCombos.Title)
		assert.Equal(t, comboOutputs[0].Status, modelCombos.Status)
	}
}
