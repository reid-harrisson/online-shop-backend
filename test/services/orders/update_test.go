package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	ordsvc "OnlineStoreBackend/services/orders"

	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
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
