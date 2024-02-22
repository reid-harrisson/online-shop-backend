package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cartsvc "OnlineStoreBackend/services/cart_items"
	etsvc "OnlineStoreBackend/services/email_templates"
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
// @Success 201 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order [post]
func (h *HandlersOrderManagement) Create(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	modelCarts := make([]models.CartItemsWithDetail, 0)
	modelTax := models.TaxSettings{}
	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	taxRepo.ReadTaxSetting(&modelTax, customerID)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadDetail(&modelCarts, customerID)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.DeleteAll(customerID)

	modelOrder := models.Orders{}
	orderService := ordsvc.NewServiceOrder(h.server.DB)
	orderService.Create(&modelOrder, modelCarts, modelTax, customerID)
	modelItems := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelItems, uint64(modelOrder.ID))
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

	orderService := ordsvc.NewServiceOrder(h.server.DB)
	orderService.UpdateStatus(storeID, orderID, status)

	mailData := utils.MailData{
		Name:            "PockitTV Contact Centre",
		EmailFrom:       "araki@pockittv.com",
		EmailTo:         "kaspersky3550879@gmail.com",
		EmailPretext:    "Contact Centre",
		Company:         "PockitTV",
		Subject:         "Account Activation",
		Phone:           "+12387621342",
		SourceChannel:   "Sports",
		BodyBlock:       "Body Block",
		TargetTeam:      "PockitTv Contact Team",
		BodyCtaBtnLabel: "ACTIVATE",
		// BodyCtaBtnLink:             tempUser.ActivationLink,
		BodyGreeting: "Hi",
		BodyHeading:  "ACTIVATE YOUR ACCOUNT",
		CompanyID:    2,
		// FirstName:                  tempUser.FirstName,
		HeaderPosterImageUrl:       "",
		HeaderPosterSloganSubtitle: "Activate your world of online streaming right now.",
		HeaderPosterSloganTitle:    "ARE YOU READY?",
	}
	utils.HelperMail(h.server.Config.ExternalURL.String(), c, mailData)

	modelOrder := models.StoreOrders{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByStoreAndOrderID(&modelOrder, orderID, storeID)

	return responses.NewResponseStoreOrder(c, http.StatusOK, modelOrder)
}

// Refresh godoc
// @Summary Edit order billing address
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Order ID"
// @Param address_id query int true "Address ID"
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/billing-address/{id} [put]
func (h *HandlersOrderManagement) UpdateBillingAddress(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	addressID, _ := strconv.ParseUint(c.QueryParam("address_id"), 10, 64)

	orderService := ordsvc.NewServiceOrder(h.server.DB)
	orderService.UpdateBillingAddress(orderID, addressID)

	modelOrder := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelOrder, orderID)

	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusOK, modelOrder)
}

// Refresh godoc
// @Summary Edit order shipping address
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Order ID"
// @Param address_id query int true "Address ID"
// @Success 200 {object} responses.ResponseCustomerOrderWithDetail
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/shipping-address/{id} [put]
func (h *HandlersOrderManagement) UpdateShippingAddress(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	addressID, _ := strconv.ParseUint(c.QueryParam("address_id"), 10, 64)

	orderService := ordsvc.NewServiceOrder(h.server.DB)
	orderService.UpdateShippingAddress(orderID, addressID)

	modelOrder := models.CustomerOrdersWithAddress{}
	orderRepo := repositories.NewRepositoryOrder(h.server.DB)
	orderRepo.ReadByOrderID(&modelOrder, orderID)

	return responses.NewResponseCustomerOrdersWithDetail(c, http.StatusOK, modelOrder)
}

// Refresh godoc
// @Summary Read Email Templates By Store ID
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestEmailTemplate true "Email Template Data"
// @Success 200 {object} responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/email-template [post]
func (h *HandlersOrderManagement) CreateEmailTemplate(c echo.Context) error {
	requestEmailTemplate := new(requests.RequestEmailTemplate)

	if err := c.Bind(requestEmailTemplate); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestEmailTemplate.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelEmailTemplate := models.EmailTemplate{}
	prodService := etsvc.NewServiceEmailTemplate(h.server.DB)
	prodService.Create(&modelEmailTemplate, requestEmailTemplate)

	return responses.NewResponseEmailTemplate(c, http.StatusCreated, &modelEmailTemplate)
}

// Refresh godoc
// @Summary Read Email Templates By Store ID
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} []responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/email-template/{id} [get]
func (h *HandlersOrderManagement) ReadEmailTemplateByStoreID(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelEmailTemplates := make([]models.EmailTemplate, 0)

	emailTemplateRepository := repositories.NewRepositoryEmailTemplate(h.server.DB)
	emailTemplateRepository.ReadEmailTemplateByStoreID(&modelEmailTemplates, storeID)

	return responses.NewResponseEmailTemplates(c, http.StatusOK, modelEmailTemplates)
}

// Refresh godoc
// @Summary Update Email Template
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Email Template ID"
// @Param params body requests.RequestEmailTemplate true "Email Template Data"
// @Success 200 {object} responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/email-template/{id} [put]
func (h *HandlersOrderManagement) UpdateEmailTemplate(c echo.Context) error {
	emailTemplateID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	requestEmailTemplate := new(requests.RequestEmailTemplate)

	if err := c.Bind(requestEmailTemplate); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := requestEmailTemplate.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelEmailTemplate := models.EmailTemplate{}

	emailTemplateService := etsvc.NewServiceEmailTemplate(h.server.DB)
	if err := emailTemplateService.Update(emailTemplateID, &modelEmailTemplate, requestEmailTemplate); err != nil {
		return responses.Response(c, http.StatusBadRequest, "No record found.")
	}

	return responses.NewResponseEmailTemplate(c, http.StatusCreated, &modelEmailTemplate)
}

// Refresh godoc
// @Summary Delete Email Template
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Email Template ID"
// @Success 200 {object} []responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/order/email-template/{id} [delete]
func (h *HandlersOrderManagement) DeleteEmailTemplate(c echo.Context) error {
	emailTemplateID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	emailTemplateService := etsvc.NewServiceEmailTemplate(h.server.DB)
	if err := emailTemplateService.Delete(emailTemplateID); err != nil {
		return responses.Response(c, http.StatusBadRequest, "No record found.")
	}

	return responses.MessageResponse(c, http.StatusCreated, "Successfully deleted")
}
