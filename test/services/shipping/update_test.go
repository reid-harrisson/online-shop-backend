package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateShipping(t *testing.T) {
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
	if assert.NoError(t, shipService.Update(1, &reqShippingData, &modelShipData)) {
		assert.Equal(t, reqShippingData.Weight, modelShipData.Weight)
		assert.Equal(t, reqShippingData.Width, modelShipData.Width)
		assert.Equal(t, reqShippingData.Height, modelShipData.Height)
		assert.Equal(t, reqShippingData.Length, modelShipData.Length)
	}
}
