package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
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

type HandlersOrder struct {
	server *s.Server
}

func NewHandlersOrders(server *s.Server) *HandlersOrder {
	return &HandlersOrder{server: server}
}

// Refresh godoc
// @Summary Add order
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 201 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order [post]
func (h *HandlersOrder) Create(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCheckout)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err = couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCarts := make([]models.CartItemsWithDetail, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadDetail(&modelCarts, customerID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	err = cartService.DeleteAll(customerID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelOrder := models.Orders{}
	modelItems := []models.OrderItems{}
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	err = ordService.Create(&modelOrder, &modelItems, modelCarts, req.BillingAddressID, req.ShippingAddressID, modelCoupons, customerID, models.Combos{})
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	totalAmount := 0.0

	for _, modelItem := range modelItems {
		totalAmount += modelItem.TotalPrice
	}

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

	return responses.NewResponseOrderItems(c, http.StatusCreated, modelItems)
}

// Refresh godoc
// @Summary Add order with combo
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param combo_id query int true "Combo ID"
// @Param params body requests.RequestCheckout true "Address and coupon"
// @Success 201 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/combo [post]
func (h *HandlersOrder) CreateCombo(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	comboID, err := strconv.ParseUint(c.QueryParam("combo_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCheckout)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCombo := models.Combos{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadByID(&modelCombo, comboID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err = couRepo.ReadByIDs(&modelCoupons, req.CouponIDs)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelCarts := make([]models.CartItemsWithDetail, 0)
	err = combRepo.ReadDetail(&modelCarts, comboID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelOrder := models.Orders{}
	modelItems := []models.OrderItems{}
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	err = ordService.Create(&modelOrder, &modelItems, modelCarts, req.BillingAddressID, req.ShippingAddressID, modelCoupons, customerID, modelCombo)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	totalAmount := 0.0

	for _, modelItem := range modelItems {
		totalAmount += modelItem.TotalPrice
	}

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

	return responses.NewResponseOrderItems(c, http.StatusCreated, modelItems)
}

// Refresh godoc
// @Summary Read orders by ID
// @Tags Order Actions
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/{id} [get]
func (h *HandlersOrder) ReadByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelOrders := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	err = orderRepo.ReadByOrderID(&modelOrders, id)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Read orders by Store
// @Tags Order Actions
// @Accept json
// @Produce json
// @Param store_id query int false "Store ID"
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/store [get]
func (h *HandlersOrder) ReadByStoreID(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelOrders := make([]models.StoreOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	err = orderRepo.ReadByStoreID(&modelOrders, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseStoreOrders(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Read orders by Customer
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/customer [get]
func (h *HandlersOrder) ReadByCustomerID(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelOrders := make([]models.CustomerOrders, 0)
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	err = orderRepo.ReadByCustomerID(&modelOrders, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCustomerOrders(c, http.StatusOK, modelOrders)
}

// Refresh godoc
// @Summary Edit order status
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Order ID"
// @Param store_id query int true "Store ID"
// @Param status query string ture "Status"
// @Success 200 {object} responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/status/{id} [put]
func (h *HandlersOrder) UpdateStatus(c echo.Context) error {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := c.QueryParam("status")

	modelItems := make([]models.OrderItems, 0)
	ordService := ordsvc.NewServiceOrder(h.server.DB)
	err = ordService.UpdateStatus(&modelItems, storeID, orderID, status)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// mailData := utils.MailData{
	// 	Name:                       "PockitTV Contact Centre",
	// 	EmailFrom:                  "araki@pockittv.com",
	// 	EmailTo:                    "kaspersky3550879@gmail.com",
	// 	EmailPretext:               "Contact Centre",
	// 	Company:                    "PockitTV",
	// 	Subject:                    "Account Activation",
	// 	Phone:                      "+12387621342",
	// 	SourceChannel:              "Sports",
	// 	BodyBlock:                  "Body Block",
	// 	TargetTeam:                 "PockitTv Contact Team",
	// 	BodyCtaBtnLabel:            "ACTIVATE",
	// 	BodyCtaBtnLink:             "",
	// 	BodyGreeting:               "Hi",
	// 	BodyHeading:                "ACTIVATE YOUR ACCOUNT",
	// 	CompanyID:                  2,
	// 	FirstName:                  "",
	// 	HeaderPosterImageUrl:       "",
	// 	HeaderPosterSloganSubtitle: "Activate your world of online streaming right now.",
	// 	HeaderPosterSloganTitle:    "ARE YOU READY?",
	// }
	// utils.HelperMail(h.server.Config.Services.CommonTool, c, mailData)

	return responses.NewResponseOrderItems(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Update order item status
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id query int true "Order Item ID"
// @Param status query string ture "Status"
// @Success 200 {object} responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/status [put]
func (h *HandlersOrder) UpdateOrderItemStatus(c echo.Context) error {
	orderID, err := strconv.ParseUint(c.QueryParam("order_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := c.QueryParam("status")

	orderService := ordsvc.NewServiceOrder(h.server.DB)
	err = orderService.UpdateOrderItemStatus(orderID, status)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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

// Refresh godoc
// @Summary Generate PDF
// @Tags Order Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id query int true "Order Item ID"
// @Success 200 {object} responses.ResponseStoreOrder
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/order/generate-pdf [put]
func (h *HandlersOrder) GeneratePDF(c echo.Context) error {
	_, err := strconv.ParseUint(c.QueryParam("order_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// orderService := ordsvc.NewServiceOrder(h.server.DB)
	// err = orderService.GeneratePDF(orderID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusCreated, "Successfly created pdf file.")
}
