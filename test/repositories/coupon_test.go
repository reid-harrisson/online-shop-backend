package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	couponOutputs = []models.Coupons{
		{
			StoreID:           1,
			CouponCode:        "30PERCENTMIN",
			DiscountType:      0,
			CouponAmount:      30.00,
			AllowFreeShipping: 0,
			MinimumSpend:      30,
			MaximumSpend:      100,
		},
	}
)

func TestCouponReadByStoreID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCouponsDB(db)

	// Setup
	var couponRepo = repositories.NewRepositoryCoupon(db)

	var modelCoupons = []models.Coupons{}

	// Assertions
	if assert.NoError(t, couponRepo.ReadByStoreID(&modelCoupons, 1)) {
		couponOutputs[0].ID = modelCoupons[0].ID
		couponOutputs[0].CreatedAt = modelCoupons[0].CreatedAt
		couponOutputs[0].UpdatedAt = modelCoupons[0].UpdatedAt
		couponOutputs[0].ExpiryDate = modelCoupons[0].ExpiryDate

		assert.Equal(t, couponOutputs[0], modelCoupons[0])
	}
}

func TestCouponReadByID(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCouponsDB(db)

	// Setup
	var couponRepo = repositories.NewRepositoryCoupon(db)

	var modelCoupon = models.Coupons{}

	// Assertions
	if assert.NoError(t, couponRepo.ReadByID(&modelCoupon, 1)) {
		couponOutputs[0].ID = modelCoupon.ID
		couponOutputs[0].CreatedAt = modelCoupon.CreatedAt
		couponOutputs[0].UpdatedAt = modelCoupon.UpdatedAt
		couponOutputs[0].ExpiryDate = modelCoupon.ExpiryDate

		assert.Equal(t, couponOutputs[0], modelCoupon)
	}
}

func TestCouponReadByCode(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCouponsDB(db)

	// Setup
	var couponRepo = repositories.NewRepositoryCoupon(db)

	var modelCoupon = models.Coupons{}

	// Assertions
	var code = "30PERCENTMIN"
	if assert.NoError(t, couponRepo.ReadByCode(&modelCoupon, code)) {
		couponOutputs[0].ID = modelCoupon.ID
		couponOutputs[0].CreatedAt = modelCoupon.CreatedAt
		couponOutputs[0].UpdatedAt = modelCoupon.UpdatedAt
		couponOutputs[0].ExpiryDate = modelCoupon.ExpiryDate

		assert.Equal(t, couponOutputs[0], modelCoupon)
	}
}

func TestCouponReadByIDs(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetCouponsDB(db)

	// Setup
	var couponRepo = repositories.NewRepositoryCoupon(db)

	var modelCoupons = []models.Coupons{}
	var ids = []uint64{1}

	// Assertions
	if assert.NoError(t, couponRepo.ReadByIDs(&modelCoupons, ids)) {
		couponOutputs[0].ID = modelCoupons[0].ID
		couponOutputs[0].CreatedAt = modelCoupons[0].CreatedAt
		couponOutputs[0].UpdatedAt = modelCoupons[0].UpdatedAt
		couponOutputs[0].ExpiryDate = modelCoupons[0].ExpiryDate

		assert.Equal(t, couponOutputs[0], modelCoupons[0])
	}
}
