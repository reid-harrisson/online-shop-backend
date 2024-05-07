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
			Price:       76,
			Quantity:    1,
			TotalPrice:  76,
		},
	}
	readCustomerInsightsOutputs = models.CustomerInsights{
		MaleCount:   1,
		FemaleCount: 0,
		AverageAge:  39,
		YoungestAge: 39,
		OldestAge:   39,
		Location:    "",
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
			assert.Equal(t, readSaleReportOutputs[0], modelReports[0])
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
		assert.Equal(t, readCustomerInsightsOutputs, modelReports)
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
		assert.Equal(t, readConvenRateOutputs, modelReports)
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

func TestAnalyticsReadCLV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = []models.CustomerSales{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadCLV(&modelTimes, 1, startDate, endDate))
}

func TestAnalyticsReadRevenue(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = models.StoreSales{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadRevenue(&modelTimes, 1, startDate, endDate))
}

func TestAnalyticsReadAOV(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetVisitorsDB(db)

	// Setup
	analRepo := repositories.NewRepositoryAnalytics(db)
	var modelTimes = models.StoreSales{}

	// Assert
	var startDate = time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)
	var endDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	assert.NoError(t, analRepo.ReadAOV(&modelTimes, 1, startDate, endDate))
}
