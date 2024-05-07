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
				ClassID:     0,
				Condition:   0,
				Min:         0,
				Max:         0,
				Break:       0,
				Abort:       0,
				RowCost:     2,
				ItemCost:    1,
				CostPerKg:   0.2,
				PercentCost: 1,
			},
		},
		2: {
			{
				MethodID:    1,
				ClassID:     0,
				Condition:   0,
				Min:         0,
				Max:         0,
				Break:       0,
				Abort:       0,
				RowCost:     3,
				ItemCost:    2,
				CostPerKg:   0.1,
				PercentCost: 1.1,
			},
		}}
	readMapMeth = map[uint64]models.ShippingMethods{
		1: {
			StoreID: 1,
		},
		2: {
			StoreID: 2,
		},
	}
)

func TestReadMethodAndTableRatesByStoreIDsShippingMethod(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetShippingData(db)
	test_utils.ResetShippingMethodsDB(db)
	test_utils.ResetShippingTableRatesDB(db)

	// Setup
	shipRepo := repositories.NewRepositoryShippingMethod(db)

	mapRates := make(map[uint64][]models.ShippingTableRates, 0)
	mapMeth := make(map[uint64]models.ShippingMethods, 0)

	// Assertions
	if assert.NoError(t, shipRepo.ReadMethodAndTableRatesByStoreIDs(&mapRates, &mapMeth, []uint64{1, 2})) {
		if assert.Equal(t, len(readMapRates), len(mapRates)) {
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
}
