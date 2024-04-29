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
