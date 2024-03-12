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

type HandlersCheckoutProcess struct {
	server *s.Server
}

func NewHandlersCheckoutProcess(server *s.Server) *HandlersCheckoutProcess {
	return &HandlersCheckoutProcess{server: server}
}

// Refresh godoc
// @Summary Read checkout
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseCheckout
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout [get]
func (h *HandlersCheckoutProcess) Read(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo.ReadDetail(&modelItems, customerID)
	return responses.NewResponseCart(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Create address to customer
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestAddress true "Address"
// @Success 201 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address [post]
func (h *HandlersCheckoutProcess) CreateAddress(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	req := new(requests.RequestAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.Addresses{}
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	addrService.Create(&modelAddr, req, customerID)

	return responses.NewResponseAddress(c, http.StatusCreated, modelAddr)
}

// Refresh godoc
// @Summary Read addresses
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address [get]
func (h *HandlersCheckoutProcess) ReadAddresses(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelAddrs := make([]models.Addresses, 0)
	addrRepo := repositories.NewRepositoryAddresses(h.server.DB)
	addrRepo.ReadAddressesByCustomerID(&modelAddrs, customerID)

	return responses.NewResponseAddresses(c, http.StatusOK, modelAddrs)
}

// Refresh godoc
// @Summary Read coupon
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param code query string true "Customer ID"
// @Success 200 {object} []responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/coupon [get]
func (h *HandlersCheckoutProcess) ReadCoupon(c echo.Context) error {
	code := c.QueryParam("code")

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByCode(&modelCoupon, code)

	return responses.NewResponseCoupon(c, http.StatusOK, modelCoupon)
}

// Refresh godoc
// @Summary Update address to customer
// @Tags Customer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Address ID"
// @Param params body requests.RequestAddress true "Customer Address"
// @Success 200 {object} []responses.ResponseAddress
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout/address/{id} [put]
func (h *HandlersCheckoutProcess) UpdateAddress(c echo.Context) error {
	addressID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestAddress)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelAddr := models.Addresses{}
	addrService := addrsvc.NewServiceAddress(h.server.DB)
	addrService.Update(&modelAddr, req, addressID)

	return responses.NewResponseAddress(c, http.StatusOK, modelAddr)
}

// Refresh godoc
// @Summary Update checkout
// @Tags Checkout Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.RequestCheckout
// @Success 200 {object} []responses.ResponseCheckout
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/checkout [put]
func (h *HandlersCheckoutProcess) Update(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	req := new(requests.RequestCheckout)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo.ReadDetail(&modelItems, customerID)
	return responses.NewResponseCart(c, http.StatusOK, modelItems)
}
