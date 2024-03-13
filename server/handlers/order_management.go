package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cartsvc "OnlineStoreBackend/services/cart_items"
	ordsvc "OnlineStoreBackend/services/orders"
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
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 201 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order [post]
func (h *HandlersOrderManagement) Create(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	req := new(requests.RequestCheckout)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := req.RequestCheckoutValidate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty!")
	}

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)

	modelCarts := make([]models.CartItemsWithDetail, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadDetail(&modelCarts, customerID)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.DeleteAll(customerID)

	modelOrder := models.Orders{}
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	ordService.Create(&modelOrder, modelCarts, req.BillingAddressID, req.ShippingAddressID, modelCoupons, customerID, models.Combos{})

	modelItems := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelItems, uint64(modelOrder.ID))

	var totalAmount float64

	orderRepo.CalcTotalAmount(&totalAmount, modelOrder.ID)

	currency := "usd"

	invokeData := utils.InvokeData{
		CardNumber:  req.CardNumber,
		ExpMonth:    req.ExpMonth,
		ExpYear:     req.ExpYear,
		CVC:         req.CVC,
		Amount:      totalAmount,
		Currency:    currency,
		PaymentType: utils.StorePurchase,
		RequestID:   uint64(modelOrder.ID),
	}
	utils.HelperInvoke("POST", h.server.Config.Services.TransactionServer+"/card-payment", c, invokeData)

	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusCreated, modelItems)
}

// Refresh godoc
// @Summary Add order with combo
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param combo_id query int true "Combo ID"
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 201 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/combo [post]
func (h *HandlersOrderManagement) CreateCombo(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	comboID, _ := strconv.ParseUint(c.QueryParam("combo_id"), 10, 64)

	req := new(requests.RequestCheckout)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := req.RequestCheckoutValidate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty!")
	}

	modelCombo := models.Combos{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	combRepo.ReadByID(&modelCombo, comboID)

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)

	modelCarts := make([]models.CartItemsWithDetail, 0)
	combRepo.ReadDetail(&modelCarts, comboID)

	modelOrder := models.Orders{}
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	ordService.Create(&modelOrder, modelCarts, req.BillingAddressID, req.ShippingAddressID, modelCoupons, customerID, modelCombo)

	modelItems := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelItems, uint64(modelOrder.ID))

	var totalAmount float64

	orderRepo.CalcTotalAmount(&totalAmount, modelOrder.ID)

	currency := "usd"

	invokeData := utils.InvokeData{
		CardNumber:  req.CardNumber,
		ExpMonth:    req.ExpMonth,
		ExpYear:     req.ExpYear,
		CVC:         req.CVC,
		Amount:      totalAmount,
		Currency:    currency,
		PaymentType: utils.StorePurchase,
		RequestID:   uint64(modelOrder.ID),
	}
	utils.HelperInvoke("POST", h.server.Config.Services.TransactionServer+"/card-payment", c, invokeData)

	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusCreated, modelItems)
}

// Refresh godoc
// @Summary Read orders by ID
// @Tags Order Management
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/{id} [get]
func (h *HandlersOrderManagement) ReadByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelOrders := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelOrders, id)
	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Read orders by Store
// @Tags Order Management
// @Accept json
// @Produce json
// @Param store_id query int false "Store ID"
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/store [get]
func (h *HandlersOrderManagement) ReadByStoreID(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelOrders := make([]models.StoreOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByStoreID(&modelOrders, storeID)
	return responses.NewResponseStoreOrders(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Read orders by Customer
// @Tags Order Management
// @Accept json
// @Produce json
// @Param customer_id query int false "Customer ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/customer [get]
func (h *HandlersOrderManagement) ReadByCustomerID(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelOrders := make([]models.CustomerOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByCustomerID(&modelOrders, customerID)
	return responses.NewResponseCustomerOrders(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Edit order status
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Order ID"
// @Param store_id query int true "Store ID"
// @Param status query string ture "Status"
// @Success 200 {object} responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/status/{id} [put]
func (h *HandlersOrderManagement) UpdateStatus(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	status := c.QueryParam("status")

	modelItems := make([]models.OrderItems, 0)
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	ordService.UpdateStatus(&modelItems, storeID, orderID, status)

	mailData := utils.MailData{
		Name:                       "PockitTV Contact Centre",
		EmailFrom:                  "araki@pockittv.com",
		EmailTo:                    "kaspersky3550879@gmail.com",
		EmailPretext:               "Contact Centre",
		Company:                    "PockitTV",
		Subject:                    "Account Activation",
		Phone:                      "+12387621342",
		SourceChannel:              "Sports",
		BodyBlock:                  "Body Block",
		TargetTeam:                 "PockitTv Contact Team",
		BodyCtaBtnLabel:            "ACTIVATE",
		BodyCtaBtnLink:             "",
		BodyGreeting:               "Hi",
		BodyHeading:                "ACTIVATE YOUR ACCOUNT",
		CompanyID:                  2,
		FirstName:                  "",
		HeaderPosterImageUrl:       "",
		HeaderPosterSloganSubtitle: "Activate your world of online streaming right now.",
		HeaderPosterSloganTitle:    "ARE YOU READY?",
	}
	utils.HelperMail(h.server.Config.Services.CommonTool, c, mailData)

	return responses.NewResponseOrderItems(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Update order item status
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id query int true "Order Item ID"
// @Param status query string ture "Status"
// @Success 200 {object} responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/status [put]
func (h *HandlersOrderManagement) UpdateOrderItemStatus(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.QueryParam("order_id"), 10, 64)
	status := c.QueryParam("status")

	orderService := ordsvc.NewServiceOrder(h.server.DB)
	orderService.UpdateOrderItemStatus(orderID, status)

	mailData := utils.MailData{
		Name:                       "PockitTV Contact Centre",
		EmailFrom:                  "araki@pockittv.com",
		EmailTo:                    "kaspersky3550879@gmail.com",
		EmailPretext:               "Contact Centre",
		Company:                    "PockitTV",
		Subject:                    "Account Activation",
		Phone:                      "+12387621342",
		SourceChannel:              "Sports",
		BodyBlock:                  "Body Block",
		TargetTeam:                 "PockitTv Contact Team",
		BodyCtaBtnLabel:            "ACTIVATE",
		BodyCtaBtnLink:             "",
		BodyGreeting:               "Hi",
		BodyHeading:                "ACTIVATE YOUR ACCOUNT",
		CompanyID:                  2,
		FirstName:                  "",
		HeaderPosterImageUrl:       "",
		HeaderPosterSloganSubtitle: "Activate your world of online streaming right now.",
		HeaderPosterSloganTitle:    "ARE YOU READY?",
	}
	utils.HelperMail(h.server.Config.Services.CommonTool, c, mailData)

	return responses.MessageResponse(c, http.StatusAccepted, "Order status just updated")
}
