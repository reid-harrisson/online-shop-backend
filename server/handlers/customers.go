package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	addrsvc "OnlineStoreBackend/services/customer_addresses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCustomers struct {
	server *s.Server
}

func NewHandlersCustomers(server *s.Server) *HandlersCustomers {
	return &HandlersCustomers{server: server}
}

// Refresh godoc
// @Summary Add address to customer
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestCustomerAddress true "Customer Address"
// @Success 201 {object} []responses.ResponseCustomerAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/customer/address [post]
func (h *HandlersCustomers) CreateCustomerAddress(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	req := new(requests.RequestCustomerAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.CustomerAddresses{}
	addrService := addrsvc.NewServiceCustomerAddress(h.server.DB)
	addrService.Create(&modelAddr, req, customerID)

	return responses.NewResponseCustomerAddress(c, http.StatusCreated, modelAddr)
}

// Refresh godoc
// @Summary Read address to customer
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseCustomerAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/customer/address [get]
func (h *HandlersCustomers) ReadCustomerAddress(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelAddrs := make([]models.CustomerAddresses, 0)
	cusRepo := repositories.NewRepositoryCustomer(h.server.DB)
	cusRepo.ReadAddressByCustomerID(&modelAddrs, customerID)

	return responses.NewResponseCustomerAddresses(c, http.StatusOK, modelAddrs)
}

// Refresh godoc
// @Summary Update address to customer
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Address ID"
// @Param params body requests.RequestCustomerAddress true "Customer Address"
// @Success 201 {object} []responses.ResponseCustomerAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/customer/address/{id} [put]
func (h *HandlersCustomers) UpdateCustomerAddress(c echo.Context) error {
	addressID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestCustomerAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.CustomerAddresses{}
	addrService := addrsvc.NewServiceCustomerAddress(h.server.DB)
	addrService.Update(&modelAddr, req, addressID)

	return responses.NewResponseCustomerAddress(c, http.StatusCreated, modelAddr)
}
