package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	readInventories = []models.Inventories{
		{
			ProductID:         1,
			ProductName:       "Gochujang - Korean Chilli Pepper Paste",
			VariationID:       1,
			VariationName:     "Gochujang - Korean Chilli Pepper Paste - 125G",
			StockLevel:        10,
			MinimumStockLevel: 0,
		},
	}
	readVars = []models.Variations{
		{
			Model: gorm.Model{
				ID: 1,
			},
			ProductID:       1,
			Sku:             "44-125G",
			Price:           96,
			StockLevel:      10,
			DiscountAmount:  20,
			DiscountType:    1,
			ImageUrls:       "",
			Description:     "Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.",
			Title:           "Gochujang - Korean Chilli Pepper Paste - 125G",
			BackOrderStatus: 0,
		},
	}
)

func TestReadInventoriesInventoryManagement(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVariationsDB(db)

	// Setup
	modelInventories := []models.Inventories{}
	invenRepo := repositories.NewRepositoryInventory(db)

	// Assertions
	if assert.NoError(t, invenRepo.ReadInventories(&modelInventories, 1)) {
		assert.Equal(t, readInventories, modelInventories)
	}
}

func TestReadByIDInventoryManagement(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVariationsDB(db)

	// Setup
	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(db)

	// Assertions
	if assert.NoError(t, varRepo.ReadByID(&modelVar, 1)) {
		readVars[0].Model.ID = modelVar.Model.ID
		readVars[0].CreatedAt = modelVar.CreatedAt
		readVars[0].UpdatedAt = modelVar.UpdatedAt

		assert.Equal(t, readVars[0], modelVar)
	}
}

func TestGetMinimumStockLevelInventoryManagement(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	productRepo := repositories.NewRepositoryProduct(db)

	minimumStockLevel := float64(0)

	// Assertions
	if assert.NoError(t, productRepo.GetMinimumStockLevel(&minimumStockLevel, 1)) {
		assert.Equal(t, float64(0), minimumStockLevel)
	}
}
