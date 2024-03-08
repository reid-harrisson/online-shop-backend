package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	orditmsvc "OnlineStoreBackend/services/order_items"
	classsvc "OnlineStoreBackend/services/shipping_classes"
	methodsvc "OnlineStoreBackend/services/shipping_methods"
	zonesvc "OnlineStoreBackend/services/shipping_zones"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersShippingOptions struct {
	server *s.Server
}

func NewHandlersShippingOptions(server *s.Server) *HandlersShippingOptions {
	return &HandlersShippingOptions{server: server}
}

// Refresh godoc
// @Summary Add shipping zone to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingZone true "Zone Info"
// @Success 201 {object} responses.ResponseShippingZone
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/zone [post]
func (h *HandlersShippingOptions) CreateShippingZone(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingZone)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelZone := models.ShippingZonesWithPlace{}
	zoneService := zonesvc.NewServiceShippingZone(h.server.DB)
	zoneService.Create(storeID, req, &modelZone)
	return responses.NewResponseShippingZone(c, http.StatusCreated, modelZone)
}

// Refresh godoc
// @Summary Update shipping zone
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Zone ID"
// @Param params body requests.RequestShippingZone true "Zone Info"
// @Success 201 {object} responses.ResponseShippingZone
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/zone/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingZone(c echo.Context) error {
	zoneID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingZone)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelZone := models.ShippingZonesWithPlace{}
	zoneRepo := repositories.NewRepositoryShippingZone(h.server.DB)
	zoneRepo.ReadDetailByID(&modelZone, zoneID)
	zoneService := zonesvc.NewServiceShippingZone(h.server.DB)
	zoneService.Update(req, &modelZone)
	return responses.NewResponseShippingZone(c, http.StatusOK, modelZone)
}

// Refresh godoc
// @Summary Add shipping class to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingClass true "Class Info"
// @Success 201 {object} responses.ResponseShippingClass
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/class [post]
func (h *HandlersShippingOptions) CreateShippingClass(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingClass)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelClass := models.ShippingClasses{}
	classService := classsvc.NewServiceShippingClass(h.server.DB)
	classService.Create(storeID, req, &modelClass)
	return responses.NewResponseShippingClass(c, http.StatusCreated, modelClass)
}

// Refresh godoc
// @Summary Add local pickup to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingLocalPickup true "Class Info"
// @Success 201 {object} responses.ResponseShippingLocalPickup
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/local-pickup [post]
func (h *HandlersShippingOptions) CreateShippingLocalPickup(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingLocalPickup)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	metService := methodsvc.NewServiceShippingMethod(h.server.DB)
	if err := metService.CreateShippingLocalPickup(storeID, req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingLocalPickup(c, http.StatusCreated, modelMethod)
}

// Refresh godoc
// @Summary Update local pickup method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query int path "Method ID"
// @Param params body requests.RequestShippingLocalPickup true "Class Info"
// @Success 201 {object} responses.ResponseShippingLocalPickup
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/local-pickup/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingLocalPickup(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingLocalPickup)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	methRepo.ReadByID(&modelMethod, methodID)
	if modelMethod.Method != utils.PickUp {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	methService := methodsvc.NewServiceShippingMethod(h.server.DB)
	if err := methService.UpdateShippingLocalPickup(req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingLocalPickup(c, http.StatusCreated, modelMethod)
}

// Refresh godoc
// @Summary Update shipping class
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Class ID"
// @Param params body requests.RequestShippingClass true "Class Info"
// @Success 201 {object} responses.ResponseShippingClass
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/class/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingClass(c echo.Context) error {
	classID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingClass)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelClass := models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByID(&modelClass, classID)
	classService := classsvc.NewServiceShippingClass(h.server.DB)
	classService.Update(req, &modelClass)
	return responses.NewResponseShippingClass(c, http.StatusOK, modelClass)
}

// Refresh godoc
// @Summary Add shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingMethod true "Shipping Option"
// @Success 201 {object} responses.ResponseShippingMethod
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/store [post]
func (h *HandlersShippingOptions) CreateShippingMethod(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingMethod)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	shipService := methodsvc.NewServiceShippingMethod(h.server.DB)
	if err := shipService.Create(storeID, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	modelMethods := make([]models.ShippingMethods, 0)
	shipRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	shipRepo.ReadByStoreID(&modelMethods, storeID)
	return responses.NewResponseShippingMethod(c, http.StatusCreated, modelMethods)
}

// Refresh godoc
// @Summary Read shipping method of store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseShippingMethod
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/store [get]
func (h *HandlersShippingOptions) ReadShippingOption(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	modelOptions := make([]models.ShippingMethods, 0)
	shipRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	shipRepo.ReadByStoreID(&modelOptions, storeID)
	return responses.NewResponseShippingMethod(c, http.StatusOK, modelOptions)
}

// Refresh godoc
// @Summary Update shipping method of order
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param order_id query int true "Order ID"
// @Param store_id query int true "Store ID"
// @Param method_id query int true "Shipping Method ID"
// @Success 200 {object} []responses.ResponseOrderItem
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/order [put]
func (h *HandlersShippingOptions) UpdateShippingMethod(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.QueryParam("order_id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	methodID, _ := strconv.ParseUint(c.QueryParam("method_id"), 10, 64)

	orderService := orditmsvc.NewServiceOrderItem(h.server.DB)
	modelItems := make([]models.OrderItems, 0)
	orderService.UpdateShippingMethod(&modelItems, storeID, orderID, methodID)

	return responses.NewResponseOrderItems(c, http.StatusOK, modelItems)
}
