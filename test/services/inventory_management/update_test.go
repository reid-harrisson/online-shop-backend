package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	prodsvc "OnlineStoreBackend/services/products"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readVars = []models.Variations{
		{
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

func TestUpdateMinimumStockLevel(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	prodService := prodsvc.NewServiceProduct(db)

	// Assertions
	assert.NoError(t, prodService.UpdateMinimumStockLevel(1, 10))
}

func TestUpdateStockLevel(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	prodvarService := prodvarsvc.NewServiceVariation(db)
	modelVar := readVars[0]

	// Assertions
	assert.NoError(t, prodvarService.UpdateStockLevel(&modelVar, 10))
}
