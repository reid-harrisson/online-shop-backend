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
	readStockOutputs = []models.StockAnalytics{
		{
			StockIn:  30,
			StockOut: 0,
		},
	}
	readVisitorsOutputs = []models.VisitorAnalytics{
		{
			Visitor:       1,
			UniqueVisitor: 1,
			PageView:      1,
			BounceRate:    1,
		},
	}
	readConvenRateOutputs = models.ConventionRate{
		Rate: 0,
	}
	readShoppingCartAbandonmentOutputs = models.ShoppingCartAbandonment{
		Rate: 0,
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

func TestAnalyticsReadStockAnalytic(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalytic(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadStockAnalyticWeekly(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalyticWeekly(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadStockAnalyticMonthly(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalyticMonthly(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadStockAnalyticWeekDay(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalyticWeekDay(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadStockAnalyticHour(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalyticHour(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadStockAnalyticMonth(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStockTrailsDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.StockAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadStockAnalyticMonth(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readStockOutputs[0].StockIn, modelReports[0].StockIn)
		assert.Equal(t, readStockOutputs[0].StockOut, modelReports[0].StockOut)
	}
}

func TestAnalyticsReadVisitor(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = models.VisitorAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadVisitor(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readVisitorsOutputs[0].Visitor, modelReports.Visitor)
		assert.Equal(t, readVisitorsOutputs[0].UniqueVisitor, modelReports.UniqueVisitor)
		assert.Equal(t, readVisitorsOutputs[0].PageView, modelReports.PageView)
		assert.Equal(t, readVisitorsOutputs[0].BounceRate, modelReports.BounceRate)
	}
}

func TestAnalyticsReadConventionRate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = models.ConventionRate{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadConventionRate(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readConvenRateOutputs.Rate, modelReports.Rate)
	}
}

func TestAnalyticsReadShoppingCartAbandonment(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = models.ShoppingCartAbandonment{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	if assert.NoError(t, analRepo.ReadShoppingCartAbandonment(&modelReports, 1, startDate, endDate)) {
		assert.Equal(t, readShoppingCartAbandonmentOutputs.Rate, modelReports.Rate)
	}
}

func TestAnalyticsReadCheckoutFunnelAnalytics(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.CheckoutFunnelAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadCheckoutFunnelAnalytics(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadFullFunnelAnalytics(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.FullFunnelAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadFullFunnelAnalytics(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadProductViewAnalytics(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.ProductViewAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadProductViewAnalytics(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadRepeatCustomerRate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.RepeatCustomerRates{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadRepeatCustomerRate(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadCustomerChurnRate(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = models.CustomerChurnRates{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadCustomerChurnRate(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadTopSellingProducts(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.TopSellingProducts{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	var count int
	assert.NoError(t, analRepo.ReadTopSellingProducts(&modelReports, 1, startDate, endDate, count))
}

func TestAnalyticsReadOrderTrend(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.OrderTrendAnalytics{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadOrderTrendAnalytics(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadCustomerDataByLocation(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.CustomerDataByLocation{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadCustomerDataByLocation(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadCustomerSatisfaction(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelReports = []models.CustomerSatisfaction{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadCustomerSatisfaction(&modelReports, 1, startDate, endDate))
}

func TestAnalyticsReadPageLoadingTime(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = []models.PageLoadingTime{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadPageLoadingTime(&modelTimes, 1, startDate, endDate))
}

func TestAnalyticsReadSalesByProduct(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = []models.ProductSales{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadSalesByProduct(&modelTimes, 1, startDate, endDate))
}

func TestAnalyticsReadSalesByCategory(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = []models.CategorySales{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadSalesByCategory(&modelTimes, 1, startDate, endDate))
}
