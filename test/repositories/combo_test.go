package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	comboOutputs = []models.ComboItems{
		{
			ComboID:     1,
			VariationID: 1,
			Quantity:    1,
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
		assert.Equal(t, comboOutputs[0].ComboID, modelComboItems[0].ComboID)
		assert.Equal(t, comboOutputs[0].VariationID, modelComboItems[0].VariationID)
		assert.Equal(t, comboOutputs[0].ComboID, modelComboItems[0].ComboID)
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
		assert.Equal(t, comboOutputs[0].ComboID, modelComboItems[0].ComboID)
		assert.Equal(t, comboOutputs[0].VariationID, modelComboItems[0].VariationID)
		assert.Equal(t, comboOutputs[0].ComboID, modelComboItems[0].ComboID)
	}
}
