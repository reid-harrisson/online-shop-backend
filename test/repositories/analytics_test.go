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
	readSaleReportOutputs = []models.SalesReports{
		{
			VariationID: 1,
			ProductID:   1,
			StoreID:     1,
			Price:       1,
			Quantity:    0,
			TotalPrice:  0,
		},
	}
	readCustomerInsightsOutputs = []models.CustomerInsights{
		{
			MaleCount:   0,
			FemaleCount: 0,
			AverageAge:  0,
			YoungestAge: 0,
			OldestAge:   0,
			Location:    "",
		},
	}
)

func TestAnalyticsReadSalesReport(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreOrderItemsDB(db)
	test_utils.ResetVariationsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.SalesReports{}

	// Assert
	var startDate = time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadSalesReport(&modelReports, 1, startDate, endDate)) {
		if assert.Equal(t, len(readSaleReportOutputs), len(modelReports)) {
			assert.Equal(t, readSaleReportOutputs[0].StoreID, modelReports[0].StoreID)
		}
	}
}

func TestAnalyticsReadCustomerInsights(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoreOrdersDB(db)
	test_utils.ResetStoreOrderItemsDB(db)
	test_utils.ResetUsersDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = models.CustomerInsights{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadCustomerInsights(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readCustomerInsightsOutputs[0].MaleCount, modelReports.MaleCount)
		assert.Equal(t, readCustomerInsightsOutputs[0].FemaleCount, modelReports.FemaleCount)
		assert.Equal(t, readCustomerInsightsOutputs[0].AverageAge, modelReports.AverageAge)
		assert.Equal(t, readCustomerInsightsOutputs[0].YoungestAge, modelReports.YoungestAge)
		assert.Equal(t, readCustomerInsightsOutputs[0].OldestAge, modelReports.OldestAge)
		assert.Equal(t, readCustomerInsightsOutputs[0].Location, modelReports.Location)
	}
}
