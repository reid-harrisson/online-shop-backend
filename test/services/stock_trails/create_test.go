package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/pkgs/utils"
	stocksvc "OnlineStoreBackend/services/stock_trails"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	trailInput = models.StockTrails{
		ProductID:   1,
		VariationID: 1,
		Change:      1,
		Event:       utils.ProductWarhousing,
	}
)

func TestCreateStockTrailWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetStockTrailsDB(db)

	// Setup

	var trailService = stocksvc.NewServiceStockTrail(db)
	var modelTrail = trailInput

	// Assertions
	if assert.NoError(t, trailService.Create(&modelTrail)) {
		modelTrail.CreatedAt = trailInput.CreatedAt
		modelTrail.UpdatedAt = trailInput.UpdatedAt
		modelTrail.ID = trailInput.ID
		assert.Equal(t, trailInput, modelTrail)
	}
}
