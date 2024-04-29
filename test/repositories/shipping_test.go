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
	readShipping = []models.ShippingData{
		{
			VariationID: 34,
			Weight:      5.1,
			Width:       11,
			Height:      27.5,
			Length:      16.5,
		},
	}
)

func TestReadByVariationIDShipping(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetShippingData(db)

	// Setup
	modelShipData := models.ShippingData{}
	shipRepo := repositories.NewRepositoryShippingData(db)

	// Assertions
	if assert.NoError(t, shipRepo.ReadByVariationID(&modelShipData, 34)) {
		readShipping[0].Model.ID = modelShipData.Model.ID
		readShipping[0].CreatedAt = modelShipData.CreatedAt
		readShipping[0].UpdatedAt = modelShipData.UpdatedAt

		assert.Equal(t, readShipping[0], modelShipData)
	}
}
