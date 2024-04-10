package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlersAnalytics struct {
	server *s.Server
}

func NewHandlersAnalytics(server *s.Server) *HandlersAnalytics {
	return &HandlersAnalytics{server: server}
}

// Refresh godoc
// @Summary Analyse sales reports by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseSalesReport
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales-report [get]
func (h *HandlersAnalytics) ReadSalesReports(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelReports := make([]models.SalesReports, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadSalesReport(&modelReports, storeID, startDate, endDate)
	return responses.NewResponseSalesReports(c, http.StatusOK, modelReports)
}

// Refresh godoc
// @Summary Analyse revenue
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseSalesRevenue
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/revenue [get]
func (h *HandlersAnalytics) ReadRevenue(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelSale := models.StoreSales{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadRevenue(&modelSale, storeID, startDate, endDate)
	return responses.NewResponseSalesRevenue(c, http.StatusOK, modelSale)
}

// Refresh godoc
// @Summary Analyse average order value (AOV)
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseSalesAOV
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/aov [get]
func (h *HandlersAnalytics) ReadAOV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelSale := models.StoreSales{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadAOV(&modelSale, storeID, startDate, endDate)
	return responses.NewResponseSalesRevenue(c, http.StatusOK, modelSale)
}

// Refresh godoc
// @Summary Analyse sales by product
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseSalesByProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/product [get]
func (h *HandlersAnalytics) ReadSalesByProduct(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelSales := make([]models.ProductSales, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadSalesByProduct(&modelSales, storeID, startDate, endDate)
	return responses.NewResponseSalesByProduct(c, http.StatusOK, modelSales, storeID)
}

// Refresh godoc
// @Summary Analyse sales by category
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseSalesByCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/category [get]
func (h *HandlersAnalytics) ReadSalesByCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelSales := make([]models.CategorySales, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadSalesByCategory(&modelSales, storeID, startDate, endDate)
	return responses.NewResponseSalesByCategory(c, http.StatusOK, modelSales, storeID)
}

// Refresh godoc
// @Summary Analyse customer lifetime value (CLV)
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseSalesCLV
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/clv [get]
func (h *HandlersAnalytics) ReadCLV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelSales := make([]models.CustomerSales, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCLV(&modelSales, storeID, startDate, endDate)
	return responses.NewResponseSalesCLV(c, http.StatusOK, modelSales, storeID)
}

// Refresh godoc
// @Summary Analyse customer insights by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseCustomerInsight
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/customer-insight [get]
func (h *HandlersAnalytics) ReadCustomerInsight(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelInsight := models.CustomerInsights{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerInsights(&modelInsight, storeID, startDate, endDate)
	return responses.NewResponseCustomerInsights(c, http.StatusOK, modelInsight)
}

// Refresh godoc
// @Summary Read daily stock levels by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseStockAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/stock [get]
func (h *HandlersAnalytics) ReadStockAnalytic(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelLevels := make([]models.StockAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadStockAnalytic(&modelLevels, storeID, startDate, endDate)
	return responses.NewResponseStockAnalyticsDaily(c, http.StatusOK, modelLevels)
}

// Refresh godoc
// @Summary Analyse vistors by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseVisitorAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/visitor [get]
func (h *HandlersAnalytics) ReadVisitor(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelVisitor := models.VisitorAnalytics{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadVisitor(&modelVisitor, storeID, startDate, endDate)
	return responses.NewResponseVisitorAnalytic(c, http.StatusOK, modelVisitor)
}

// Refresh godoc
// @Summary Analyse convention rate by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseConventionRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/convention-rate [get]
func (h *HandlersAnalytics) ReadConventionRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRate := models.ConventionRate{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadConventionRate(&modelRate, storeID, startDate, endDate)
	return responses.NewResponseConventionRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse shopping cart abandonment by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseShoppingCartAbandonment
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/abandonment [get]
func (h *HandlersAnalytics) ReadShoppingCartAbandonment(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRate := models.ShoppingCartAbandonment{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadShoppingCartAbandonment(&modelRate, storeID, startDate, endDate)
	return responses.NewResponseShoppingCartAbandonment(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse checkout funnel by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseCheckoutFunnelAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/checkout-funnel [get]
func (h *HandlersAnalytics) ReadCheckoutFunnelAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelFunnels := make([]models.CheckoutFunnelAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCheckoutFunnelAnalytics(&modelFunnels, storeID, startDate, endDate)
	return responses.NewResponseCheckoutFunnelAnalytics(c, http.StatusOK, modelFunnels)
}

// Refresh godoc
// @Summary Analyse full funnel by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseFullFunnelAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/full-funnel [get]
func (h *HandlersAnalytics) ReadFullFunnelAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelFunnels := make([]models.FullFunnelAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadFullFunnelAnalytics(&modelFunnels, storeID, startDate, endDate)
	return responses.NewResponseFullFunnelAnalytics(c, http.StatusOK, modelFunnels)
}

// Refresh godoc
// @Summary Analyse product views by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseProductViewAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/product-view [get]
func (h *HandlersAnalytics) ReadProductViewAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelViews := make([]models.ProductViewAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadProductViewAnalytics(&modelViews, storeID, startDate, endDate)
	return responses.NewResponseProductViewAnalytics(c, http.StatusOK, modelViews)
}

// Refresh godoc
// @Summary Analyse repeat customer rate by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseRepeatCustomerRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/repeat-rate [get]
func (h *HandlersAnalytics) ReadRepeatCustomerRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRates := make([]models.RepeatCustomerRates, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadRepeatCustomerRate(&modelRates, storeID, startDate, endDate)
	return responses.NewResponseRepeatCustomerRate(c, http.StatusOK, modelRates)
}

// Refresh godoc
// @Summary Analyse customer churn rate by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} responses.ResponseVisitorAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/churn-rate [get]
func (h *HandlersAnalytics) ReadCustomerChurnRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRate := models.CustomerChurnRates{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerChurnRate(&modelRate, storeID, startDate, endDate)
	return responses.NewResponseCustomerChurnRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse customer churn rate by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Param count query int true "Count"
// @Success 200 {object} []responses.ResponseTopSellingProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/top-selling [get]
func (h *HandlersAnalytics) ReadTopSellingProducts(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	count, _ := strconv.ParseInt(c.QueryParam("count"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelProducts := make([]models.TopSellingProducts, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadTopSellingProducts(&modelProducts, storeID, startDate, endDate, int(count))
	return responses.NewResponseTopSellingProduct(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Analyse order trend by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseOrderTrendAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/order-trend [get]
func (h *HandlersAnalytics) ReadOrderTrendAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelTrends := make([]models.OrderTrendAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadOrderTrendAnalytics(&modelTrends, storeID, startDate, endDate)
	return responses.NewResponseOrderTrendAnalytics(c, http.StatusOK, modelTrends)
}

// Refresh godoc
// @Summary Analyse customer data by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseCustomerDataByLocation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/customer-location [get]
func (h *HandlersAnalytics) ReadCustomerDataByLocation(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelLocations := make([]models.CustomerDataByLocation, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerDataByLocation(&modelLocations, storeID, startDate, endDate)
	return responses.NewResponseCustomerDataByLocation(c, http.StatusOK, modelLocations)
}

// Refresh godoc
// @Summary Analyse customer customer satisfaction by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponseCustomerSatisfaction
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/satisfaction [get]
func (h *HandlersAnalytics) ReadCustomerSatisfaction(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRates := make([]models.CustomerSatisfaction, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerSatisfaction(&modelRates, storeID, startDate, endDate)
	return responses.NewResponseCustomerSatisfaction(c, http.StatusOK, modelRates)
}

// Refresh godoc
// @Summary Analyse page loading time by store
// @Tags Analytic Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {object} []responses.ResponsePageLoadingTime
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/loading-time [get]
func (h *HandlersAnalytics) ReadPageLoadingTime(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))
	if c.QueryParam("start_date") == "" {
		startDate = time.Time{}
	}
	if c.QueryParam("end_date") == "" {
		endDate = time.Now()
	}

	modelRates := make([]models.PageLoadingTime, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadPageLoadingTime(&modelRates, storeID, startDate, endDate)
	return responses.NewResponsePageLoadingTime(c, http.StatusOK, modelRates)
}
