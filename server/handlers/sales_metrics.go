package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersSalesMetrics struct {
	server *s.Server
}

func NewHandlersSalesMetrics(server *s.Server) *HandlersSalesMetrics {
	return &HandlersSalesMetrics{server: server}
}

// Refresh godoc
// @Summary Show Revenue
// @Tags sales performance metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseSalesRevenue
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/revenue [get]
func (h *HandlersSalesMetrics) ReadRevenue(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelSale := models.StoreSales{}
	orderRepo := repositories.NewRepositorySales(h.server.DB)
	orderRepo.ReadRevenue(&modelSale, storeID)
	return responses.NewResponseSalesRevenue(c, http.StatusOK, modelSale)
}

// Refresh godoc
// @Summary Show average order value(AOV)
// @Tags sales performance metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseSalesAOV
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/aov [get]
func (h *HandlersSalesMetrics) ReadAOV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelSale := models.StoreSales{}
	orderRepo := repositories.NewRepositorySales(h.server.DB)
	orderRepo.ReadAOV(&modelSale, storeID)
	return responses.NewResponseSalesRevenue(c, http.StatusOK, modelSale)
}

// Refresh godoc
// @Summary Show sales by product
// @Tags sales performance metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseSalesByProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/product [get]
func (h *HandlersSalesMetrics) ReadSalesByProduct(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelSales := make([]models.ProductSales, 0)
	orderRepo := repositories.NewRepositorySales(h.server.DB)
	orderRepo.ReadSalesByProduct(&modelSales, storeID)
	return responses.NewResponseSalesByProduct(c, http.StatusOK, modelSales, storeID)
}

// Refresh godoc
// @Summary Show sales by category
// @Tags sales performance metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseSalesByCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/category [get]
func (h *HandlersSalesMetrics) ReadSalesByCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelSales := make([]models.CategorySales, 0)
	orderRepo := repositories.NewRepositorySales(h.server.DB)
	orderRepo.ReadSalesByCategory(&modelSales, storeID)
	return responses.NewResponseSalesByCategory(c, http.StatusOK, modelSales, storeID)
}

// Refresh godoc
// @Summary Show customer lifetime value(CLV)
// @Tags sales performance metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseSalesCLV
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/analytic/sales/clv [get]
func (h *HandlersSalesMetrics) ReadCLV(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelSales := make([]models.CustomerSales, 0)
	orderRepo := repositories.NewRepositorySales(h.server.DB)
	orderRepo.ReadCLV(&modelSales, storeID)
	return responses.NewResponseSalesCLV(c, http.StatusOK, modelSales, storeID)
}
