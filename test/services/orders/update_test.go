package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	ordsvc "OnlineStoreBackend/services/orders"

	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readItems = []models.OrderItems{
		{
			OrderID:          1,
			StoreID:          1,
			VariationID:      1,
			Price:            1,
			Quantity:         1,
			SubTotalPrice:    1,
			TaxRate:          1,
			TaxAmount:        1,
			ShippingMethodID: 1,
			ShippingPrice:    1,
			TotalPrice:       1,
			Status:           1,
		},
	}
)

func TestUpdateStatus(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	ordService := ordsvc.NewServiceOrder(db)
	modelItems := make([]models.OrderItems, 1)

	err := ordService.UpdateStatus(&modelItems, 1, 1, "1")
	assert.NoError(t, err)

}

func TestUpdateOrderItemStatus(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)

	// Setup
	ordService := ordsvc.NewServiceOrder(db)

	err := ordService.UpdateOrderItemStatus(1, "1")
	assert.NoError(t, err)

}
