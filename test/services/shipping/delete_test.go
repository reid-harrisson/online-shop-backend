package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	shipsvc "OnlineStoreBackend/services/shipping_data"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestDeleteShippingData(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetShippingData(db)

	// Setup
	shipService := shipsvc.NewServiceShippingData(db)
	shipRepo := repositories.NewRepositoryShippingData(db)

	// Assertions
	if assert.NoError(t, shipService.Delete(1)) {
		assert.Equal(t, gorm.ErrRecordNotFound, shipRepo.ReadByVariationID(&models.ShippingData{}, 1))
	}
}
