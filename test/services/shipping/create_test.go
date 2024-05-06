package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	reqShippingData = requests.RequestShippingData{
		Weight: 5.1,
		Width:  11,
		Height: 27.5,
		Length: 16.5,
	}
	shippinDataInputs = []models.ShippingData{
		{
			VariationID: 1,
			Weight:      5.100000,
			Width:       11.000000,
			Height:      27.500000,
			Length:      16.500000,
		},
		{
			VariationID: 2,
			Weight:      5.100000,
			Width:       11.000000,
			Height:      27.500000,
			Length:      16.500000,
		},
	}
	shipVarIDs  = []uint64{1, 2}
	shipIndices = map[uint64]int{
		1: 0,
		2: 1,
	}
)

func TestCreateShipping(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetTagsDB(db)
	test_utils.ResetProductTagsDB(db)

	// Setup
	shipService := shipsvc.NewServiceShippingData(db)
	modelShipData := models.ShippingData{}

	// Assertions
	assert.NoError(t, shipService.Create(1, &reqShippingData, &modelShipData))
}

func TestCreateShippingDataWithCSV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetShippingData(db)

	// Setup
	shipServie := shipsvc.NewServiceShippingData(db)
	modelNewShips := shippinDataInputs

	// Assertions
	if assert.NoError(t, shipServie.CreateWithCSV(&modelNewShips, shipVarIDs, shipIndices)) {
		assert.Equal(t, modelNewShips[0].ID, uint(1))
		assert.Equal(t, modelNewShips[1].ID, uint(2))
	}
}
