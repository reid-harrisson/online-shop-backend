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
