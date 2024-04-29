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
			ImageUrls:      "",
			Price:          96,
			Categories:     "\"Kimchi\"",
			DiscountAmount: 20,
			DiscountType:   1,
			StockLevel:     10,
			Weight:         0,
			Width:          0,
			Height:         0,
			Length:         0,
		},
	}
)

func TestReadDetailOrder(t *testing.T) {
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
