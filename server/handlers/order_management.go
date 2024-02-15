package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodOdr "OnlineStoreBackend/services/orders"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersOrderManagement struct {
	server *s.Server
}

func NewHandlersOrderManagement(server *s.Server) *HandlersOrderManagement {
	return &HandlersOrderManagement{server: server}
}

// Refresh godoc
// @Summary Add order
// @Tags order manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 201 {object} responses.ResponseCustomerOrder
// @Failure 400 {object} responses.Error
// @Router /api/v1/order [post]
func (h *HandlersOrderManagement) Create(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelCarts := make([]models.CartItemWithPrice, 0)
	modelOrders := make([]models.ProductOrders, 0)
	modelTaxSet := models.TaxSettings{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadPreview(&modelCarts, &modelTaxSet, customerID)
	orderService := prodOdr.CreateService(h.server.DB)
	if err := orderService.Create(&modelOrders, modelCarts, modelTaxSet, customerID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductOrders(c, http.StatusCreated, modelOrders)
}

// Refresh godoc
// @Summary View orders by ID
// @Tags order manangement
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrder
// @Failure 400 {object} responses.Error
// @Router /api/v1/order/{id} [get]
func (h *HandlersOrderManagement) ReadByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	fmt.Println(id)
	modelOrders := make([]models.ProductOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByID(&modelOrders, id)
	return responses.NewResponseProductOrders(c, http.StatusCreated, modelOrders)
}

// Refresh godoc
// @Summary View orders
// @Tags order manangement
// @Accept json
// @Produce json
// @Param store_id query int false "Store ID"
// @Param customer_id query int false "Customer ID"
// @Param product_id query int false "Product ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrder
// @Failure 400 {object} responses.Error
// @Router /api/v1/order [get]
func (h *HandlersOrderManagement) Read(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelOrders := make([]models.ProductOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.Read(&modelOrders, customerID, productID, storeID)
	return responses.NewResponseProductOrders(c, http.StatusCreated, modelOrders)
}

// Refresh godoc
// @Summary Edit order status
// @Tags order manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Order ID"
// @Param params body requests.RequestProductOrderStatus true "Status"
// @Success 200 {object} []responses.ResponseProductOrder
// @Failure 400 {object} responses.Error
// @Router /api/v1/order/status/{id} [put]
func (h *HandlersOrderManagement) UpdateStatus(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductOrderStatus)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelOrders := make([]models.ProductOrders, 0)
	orderService := prodOdr.CreateService(h.server.DB)
	if err := orderService.UpdateStatus(&modelOrders, req, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseProductOrders(c, http.StatusOK, modelOrders)
}
