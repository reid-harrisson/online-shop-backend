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
	readCarts = []models.CartItemsWithDetail{
		{
			CartItems: models.CartItems{
				CustomerID:  1,
				VariationID: 1,
				Quantity:    0,
			},
			StoreID:        1,
			VariationName:  "Gochujang - Korean Chilli Pepper Paste - 125G",
			ImageUrls:      "[]",
			Price:          96,
			Categories:     "\"Kefir\"",
			DiscountAmount: 20,
			DiscountType:   1,
			StockLevel:     10,
			Weight:         5.1,
			Width:          11,
			Height:         27.5,
			Length:         16.5,
		},
	}
	readCartItemsOutputs = []models.CartItems{
		{
			CustomerID:  1,
			VariationID: 1,
			Quantity:    0,
		},
	}
)

func TestCartReadDetail(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreCartItemDB(db)

	// Setup
	cartItems := []models.CartItemsWithDetail{}
	cartRepo := repositories.NewRepositoryCart(db)

	// Assertions
	if assert.NoError(t, cartRepo.ReadDetail(&cartItems, 1)) {
		readCarts[0].Model.ID = cartItems[0].Model.ID
		readCarts[0].CreatedAt = cartItems[0].CreatedAt
		readCarts[0].UpdatedAt = cartItems[0].UpdatedAt

		assert.Equal(t, readCarts[0], cartItems[0])
	}
}

func TestCartReadByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreCartItemDB(db)

	// Setup
	cartRepo := repositories.NewRepositoryCart(db)
	var modelCartItem = models.CartItems{}

	// Assertions
	if assert.NoError(t, cartRepo.ReadByID(&modelCartItem, 1)) {
		readCartItemsOutputs[0].Model.ID = modelCartItem.Model.ID
		readCartItemsOutputs[0].CreatedAt = modelCartItem.CreatedAt
		readCartItemsOutputs[0].UpdatedAt = modelCartItem.UpdatedAt

		assert.Equal(t, readCartItemsOutputs[0], modelCartItem)
	}
}

func TestCartReadByInfo(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreCartItemDB(db)

	// Setup
	cartRepo := repositories.NewRepositoryCart(db)
	var modelCartItem = models.CartItems{}

	// Assertions
	if assert.NoError(t, cartRepo.ReadByInfo(&modelCartItem, 1, 1)) {
		readCartItemsOutputs[0].Model.ID = modelCartItem.Model.ID
		readCartItemsOutputs[0].CreatedAt = modelCartItem.CreatedAt
		readCartItemsOutputs[0].UpdatedAt = modelCartItem.UpdatedAt

		assert.Equal(t, readCartItemsOutputs[0], modelCartItem)
	}
}

func TestCartReadByCustomerID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreCartItemDB(db)

	// Setup
	cartRepo := repositories.NewRepositoryCart(db)
	var modelCartItem = []models.CartItems{}

	// Assertions
	if assert.NoError(t, cartRepo.ReadByCustomerID(&modelCartItem, 1)) {
		readCartItemsOutputs[0].Model.ID = modelCartItem[0].Model.ID
		readCartItemsOutputs[0].CreatedAt = modelCartItem[0].CreatedAt
		readCartItemsOutputs[0].UpdatedAt = modelCartItem[0].UpdatedAt

		assert.Equal(t, readCartItemsOutputs[0], modelCartItem[0])
	}
}

func TestCartReadItemCount(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreCartItemDB(db)

	// Setup
	cartRepo := repositories.NewRepositoryCart(db)
	var count int64

	// Assertions
	if assert.NoError(t, cartRepo.ReadItemCount(&count, 1)) {
		assert.Equal(t, int64(2), count)
	}
}
