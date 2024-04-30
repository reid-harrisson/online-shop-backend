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
	readMapRates = map[uint64][]models.ShippingTableRates{
		1: {
			{
				MethodID:    1,
				ClassID:     1,
				Condition:   0,
				Min:         0,
				Max:         0,
				Break:       0,
				Abort:       0,
				RowCost:     0,
				ItemCost:    0,
				CostPerKg:   0,
				PercentCost: 0,
			},
		}}
	readMapMeth = map[uint64]models.ShippingMethods{
		1: {
			Title:               "",
			Description:         "",
			ZoneID:              1,
			StoreID:             1,
			Method:              0,
			Requirement:         0,
			MinimumOrderAmount:  0,
			TaxStatus:           0,
			Cost:                0,
			TaxIncluded:         0,
			HandlingFee:         0,
			MaximumShippingCost: 0,
			CalculationType:     0,
			HandlingFeePerClass: 0,
			MinimumCostPerClass: 0,
			MaximumCostPerClass: 0,
			DiscountInMinMax:    0,
			TaxInMinMax:         0,
		},
	}
)

func TestReadMethodAndTableRatesByStoreIDsShippingMethod(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetShippingData(db)

	// Setup
	shipRepo := repositories.NewRepositoryShippingMethod(db)

	mapRates := make(map[uint64][]models.ShippingTableRates, 0)
	mapMeth := make(map[uint64]models.ShippingMethods, 0)

	// Assertions
	if assert.NoError(t, shipRepo.ReadMethodAndTableRatesByStoreIDs(&mapRates, &mapMeth, []uint64{1, 2})) {
		readMapRates[1][0].ID = mapRates[1][0].ID
		readMapRates[1][0].CreatedAt = mapRates[1][0].CreatedAt
		readMapRates[1][0].UpdatedAt = mapRates[1][0].UpdatedAt

		assert.Equal(t, readMapRates[1][0], mapRates[1][0])

		meth := readMapMeth[1]
		meth.ID = mapMeth[1].ID
		meth.CreatedAt = mapMeth[1].CreatedAt
		meth.UpdatedAt = mapMeth[1].UpdatedAt

		assert.Equal(t, meth, mapMeth[1])
	}
}
