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

	var modelCoupon = []models.Coupons{}

	// Assertions
	if assert.NoError(t, couponRepo.ReadByStoreID(&modelCoupon, 1)) {
		assert.Equal(t, couponOutputs[0].StoreID, modelCoupon[0].StoreID)
		assert.Equal(t, couponOutputs[0].CouponCode, modelCoupon[0].CouponCode)
		assert.Equal(t, couponOutputs[0].DiscountType, modelCoupon[0].DiscountType)
		assert.Equal(t, couponOutputs[0].CouponAmount, modelCoupon[0].CouponAmount)
		assert.Equal(t, couponOutputs[0].AllowFreeShipping, modelCoupon[0].AllowFreeShipping)
		assert.Equal(t, couponOutputs[0].MinimumSpend, modelCoupon[0].MinimumSpend)
		assert.Equal(t, couponOutputs[0].MaximumSpend, modelCoupon[0].MaximumSpend)
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
		assert.Equal(t, couponOutputs[0].StoreID, modelCoupon.StoreID)
		assert.Equal(t, couponOutputs[0].CouponCode, modelCoupon.CouponCode)
		assert.Equal(t, couponOutputs[0].DiscountType, modelCoupon.DiscountType)
		assert.Equal(t, couponOutputs[0].CouponAmount, modelCoupon.CouponAmount)
		assert.Equal(t, couponOutputs[0].AllowFreeShipping, modelCoupon.AllowFreeShipping)
		assert.Equal(t, couponOutputs[0].MinimumSpend, modelCoupon.MinimumSpend)
		assert.Equal(t, couponOutputs[0].MaximumSpend, modelCoupon.MaximumSpend)
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
		assert.Equal(t, couponOutputs[0].StoreID, modelCoupon.StoreID)
		assert.Equal(t, couponOutputs[0].CouponCode, modelCoupon.CouponCode)
		assert.Equal(t, couponOutputs[0].DiscountType, modelCoupon.DiscountType)
		assert.Equal(t, couponOutputs[0].CouponAmount, modelCoupon.CouponAmount)
		assert.Equal(t, couponOutputs[0].AllowFreeShipping, modelCoupon.AllowFreeShipping)
		assert.Equal(t, couponOutputs[0].MinimumSpend, modelCoupon.MinimumSpend)
		assert.Equal(t, couponOutputs[0].MaximumSpend, modelCoupon.MaximumSpend)
	}
}
