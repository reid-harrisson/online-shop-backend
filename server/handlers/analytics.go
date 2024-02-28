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
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseSalesReport
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales-report [get]
func (h *HandlersAnalytics) ReadSalesReports(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelReports := make([]models.SalesReports, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadSalesReport(&modelReports, storeID)
	return responses.NewResponseSalesReports(c, http.StatusOK, modelReports)
}

// Refresh godoc
// @Summary Analyse customer insights by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} responses.ResponseCustomerInsight
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/customer-insight [get]
func (h *HandlersAnalytics) ReadCustomerInsight(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelInsight := models.CustomerInsights{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerInsights(&modelInsight, storeID, startDate, endDate)
	return responses.NewResponseCustomerInsights(c, http.StatusOK, modelInsight)
}

// Refresh godoc
// @Summary Read stock levels by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseStockLevelAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/stock-level [get]
func (h *HandlersAnalytics) ReadStockLevels(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelLevels := make([]models.StockLevelAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadStockLevelAnalytics(&modelLevels, storeID)
	return responses.NewResponseStockLevelAnalytics(c, http.StatusOK, modelLevels)
}

// Refresh godoc
// @Summary Analyse vistors by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} responses.ResponseVisitorAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/visitor [get]
func (h *HandlersAnalytics) ReadVisitor(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelVisitor := models.VisitorAnalytics{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadVisitor(&modelVisitor, storeID, startDate, endDate)
	return responses.NewResponseVisitorAnalytic(c, http.StatusOK, modelVisitor)
}

// Refresh godoc
// @Summary Analyse convention rate by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseConventionRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/convention-rate [get]
func (h *HandlersAnalytics) ReadConventionRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelRate := models.ConventionRate{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadConventionRate(&modelRate, storeID)
	return responses.NewResponseConventionRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse shopping cart abandonment by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseShoppingCartAbandonment
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/abandonment [get]
func (h *HandlersAnalytics) ReadShoppingCartAbandonment(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelRate := models.ShoppingCartAbandonment{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadShoppingCartAbandonment(&modelRate, storeID)
	return responses.NewResponseShoppingCartAbandonment(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse checkout funnel by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseCheckoutFunnelAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/checkout-funnel [get]
func (h *HandlersAnalytics) ReadCheckoutFunnelAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelFunnels := make([]models.CheckoutFunnelAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCheckoutFunnelAnalytics(&modelFunnels, storeID)
	return responses.NewResponseCheckoutFunnelAnalytics(c, http.StatusOK, modelFunnels)
}

// Refresh godoc
// @Summary Analyse full funnel by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseFullFunnelAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/full-funnel [get]
func (h *HandlersAnalytics) ReadFullFunnelAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelFunnels := make([]models.FullFunnelAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadFullFunnelAnalytics(&modelFunnels, storeID)
	return responses.NewResponseFullFunnelAnalytics(c, http.StatusOK, modelFunnels)
}

// Refresh godoc
// @Summary Analyse product views by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseProductViewAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/product-view [get]
func (h *HandlersAnalytics) ReadProductViewAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelViews := make([]models.ProductViewAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadProductViewAnalytics(&modelViews, storeID)
	return responses.NewResponseProductViewAnalytics(c, http.StatusOK, modelViews)
}

// Refresh godoc
// @Summary Analyse repeat customer rate by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseRepeatCustomerRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/repeat-rate [get]
func (h *HandlersAnalytics) ReadRepeatCustomerRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelRates := make([]models.RepeatCustomerRates, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadRepeatCustomerRate(&modelRates, storeID)
	return responses.NewResponseRepeatCustomerRate(c, http.StatusOK, modelRates)
}

// Refresh godoc
// @Summary Analyse customer churn rate by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} responses.ResponseVisitorAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/churn-rate [get]
func (h *HandlersAnalytics) ReadCustomerChurnRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelRate := models.CustomerChurnRates{}
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerChurnRate(&modelRate, storeID, startDate, endDate)
	return responses.NewResponseCustomerChurnRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Analyse customer churn rate by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Param count query int true "Count"
// @Success 200 {object} []responses.ResponseTopSellingProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/top-selling [get]
func (h *HandlersAnalytics) ReadTopSellingProducts(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	count, _ := strconv.ParseUint(c.QueryParam("count"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelProducts := make([]models.TopSellingProducts, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadTopSellingProducts(&modelProducts, storeID, startDate, endDate, count)
	return responses.NewResponseTopSellingProduct(c, http.StatusOK, modelProducts)
}

// Refresh godoc
// @Summary Analyse order trend by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} []responses.ResponseOrderTrendAnalytic
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/order-trend [get]
func (h *HandlersAnalytics) ReadOrderTrendAnalytics(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelTrends := make([]models.OrderTrendAnalytics, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadOrderTrendAnalytics(&modelTrends, storeID, startDate, endDate)
	return responses.NewResponseOrderTrendAnalytics(c, http.StatusOK, modelTrends)
}

// Refresh godoc
// @Summary Analyse customer data by store
// @Tags Analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param start_date query string true "Start Date"
// @Param end_date query string true "End Date"
// @Success 200 {object} []responses.ResponseCustomerDataByLocation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/customer-location [get]
func (h *HandlersAnalytics) ReadCustomerDataByLocation(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, c.QueryParam("start_date"))
	endDate, _ := time.Parse(layout, c.QueryParam("end_date"))

	modelLocations := make([]models.CustomerDataByLocation, 0)
	analyRepo := repositories.NewRepositoryAnalytics(h.server.DB)
	analyRepo.ReadCustomerDataByLocation(&modelLocations, storeID, startDate, endDate)
	return responses.NewResponseCustomerDataByLocation(c, http.StatusOK, modelLocations)
}
