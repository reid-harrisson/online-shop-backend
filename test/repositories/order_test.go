package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"
	"time"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readCoupons = []models.Coupons{
		{
			StoreID:           1,
			CouponCode:        "30PERCENTMIN",
			DiscountType:      0,
			CouponAmount:      30,
			AllowFreeShipping: 0,
			ExpiryDate:        time.Time{},
			MinimumSpend:      30,
			MaximumSpend:      100,
		},
	}
	readOrders = []models.CustomerOrdersWithAddress{
		{
			BillingAddress: models.Addresses{
				AddressLine1: "",
				AddressLine2: "",
				SubUrb:       "",
				CustomerID:   0,
				CountryID:    0,
				RegionID:     0,
				CityID:       0,
				PostalCode:   "",
				Active:       0,
			},
			ShippingAddress: models.Addresses{
				AddressLine1: "",
				AddressLine2: "",
				SubUrb:       "",
				CustomerID:   0,
				CountryID:    0,
				RegionID:     0,
				CityID:       0,
				PostalCode:   "",
				Active:       0,
			},
			Items: []models.CustomerOrderItems{},
		},
	}
	readCustomerOrders = []models.CustomerOrders{
		{
			OrderID:           0,
			OrderStatus:       0,
			TotalPrice:        0,
			BillingAddressID:  0,
			ShippingAddressID: 0,
		},
	}
	readStores = []models.StoreOrders{
		{
			OrderID:           1,
			CustomerID:        0,
			VariationID:       1,
			Price:             76,
			Quantity:          1,
			SubTotalPrice:     76,
			BillingAddressID:  0,
			ShippingAddressID: 0,
			TaxRate:           10,
			TaxAmount:         7.6,
			ShippingMethodID:  1,
			ShippingPrice:     0,
			TotalPrice:        83.6,
			ProductStatus:     0,
		},
	}
)

func TestReadByIDsOrder(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCouponsDB(db)

	// Setup
	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(db)

	couponIDs := []uint64{1}

	// Assertions
	if assert.NoError(t, couRepo.ReadByIDs(&modelCoupons, couponIDs)) {
		readCoupons[0].Model.ID = modelCoupons[0].Model.ID
		readCoupons[0].CreatedAt = modelCoupons[0].CreatedAt
		readCoupons[0].UpdatedAt = modelCoupons[0].UpdatedAt
		readCoupons[0].ExpiryDate = modelCoupons[0].ExpiryDate

		assert.Equal(t, readCoupons[0], modelCoupons[0])
	}
}

func TestReadByOrderIDOrder(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreCustomerAddressesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	modelOrder := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(db)

	// Assertions
	if assert.NoError(t, orderRepo.ReadByOrderID(&modelOrder, 1)) {
		assert.Equal(t, readOrders[0], modelOrder)
	}
}

func TestReadByStoreIDOrder(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreCustomerAddressesDB(db)
	test_utils.ResetStoreOrderItemsDB(db)
	test_utils.ResetVariationsDB(db)
	test_utils.ResetProductsDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	modelStore := make([]models.StoreOrders, 1)
	orderRepo := repositories.NewRepositoryOrder(db)

	// Assertions
	if assert.NoError(t, orderRepo.ReadByStoreID(&modelStore, 1)) {
		assert.Equal(t, readStores[0], modelStore[0])
	}
}

func TestReadByCustomerIDOrder(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreCustomerAddressesDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	modelOrders := make([]models.CustomerOrders, 1)
	orderRepo := repositories.NewRepositoryOrder(db)

	// Assertions
	if assert.NoError(t, orderRepo.ReadByCustomerID(&modelOrders, 1)) {
		assert.Equal(t, readCustomerOrders[0], modelOrders[0])
	}
}
