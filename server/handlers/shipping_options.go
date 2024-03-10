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
	methsvc "OnlineStoreBackend/services/shipping_methods"
	tablesvc "OnlineStoreBackend/services/shipping_table_rates"
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
// @Summary Add local pickup method to store
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
	metService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := metService.CreateShippingLocalPickup(storeID, req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingLocalPickup(c, http.StatusCreated, modelMethod)
}

// Refresh godoc
// @Summary Add free shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingFree true "Class Info"
// @Success 201 {object} responses.ResponseShippingFree
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/free [post]
func (h *HandlersShippingOptions) CreateShippingFree(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingFree)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	metService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := metService.CreateShippingFree(storeID, req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingFree(c, http.StatusCreated, modelMethod)
}

// Refresh godoc
// @Summary Add flat rate shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingFlatRate true "Class Info"
// @Success 201 {object} responses.ResponseShippingFlatRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/flat-rate [post]
func (h *HandlersShippingOptions) CreateShippingFlatRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingFlatRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingFlatRates{}
	metService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := metService.CreateShippingFlatRate(storeID, req, &modelMethod, &modelRates); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingFlatRate(c, http.StatusCreated, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Add table rate shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestShippingTableRate true "Class Info"
// @Success 201 {object} responses.ResponseShippingTableRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/table-rate [post]
func (h *HandlersShippingOptions) CreateShippingTableRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingTableRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingTableRates{}
	metService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := metService.CreateShippingTableRate(storeID, req, &modelMethod, &modelRates); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingTableRate(c, http.StatusCreated, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Add table rate shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestTableRate true "Class Info"
// @Success 201 {object} responses.ResponseShippingTableRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/rate [post]
func (h *HandlersShippingOptions) CreateShippingRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestTableRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelRate := models.ShippingTableRates{}
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	if err := tableService.Create(storeID, req, &modelRate); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This condition already exist.")
	}
	return responses.NewResponseTableRate(c, http.StatusCreated, modelRate)
}

// Refresh godoc
// @Summary Read all shipping method of store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseShippingMethod
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping [get]
func (h *HandlersShippingOptions) ReadAllShippingMethod(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestShippingTableRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethods := []models.ShippingMethods{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	methRepo.ReadByStoreID(&modelMethods, storeID)
	return responses.NewResponseShippingMethods(c, http.StatusOK, modelMethods)
}

// Refresh godoc
// @Summary Read rates
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseTableRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/rate [get]
func (h *HandlersShippingOptions) ReadShippingRate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	modelRates := []models.ShippingTableRates{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	methRepo.ReadRates(&modelRates, storeID)
	return responses.NewResponseTableRates(c, http.StatusOK, modelRates)
}

// Refresh godoc
// @Summary Read local pickup method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Success 200 {object} responses.ResponseShippingLocalPickup
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/local-pickup/{id} [get]
func (h *HandlersShippingOptions) ReadShippingLocalPickup(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelMethod := models.ShippingMethods{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadByID(&modelMethod, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.PickUp {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	return responses.NewResponseShippingLocalPickup(c, http.StatusOK, modelMethod)
}

// Refresh godoc
// @Summary Read free shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Success 200 {object} responses.ResponseShippingFree
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/free/{id} [get]
func (h *HandlersShippingOptions) ReadShippingFree(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelMethod := models.ShippingMethods{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadByID(&modelMethod, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.FreeShipping {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not free shipping method.")
	}
	return responses.NewResponseShippingFree(c, http.StatusOK, modelMethod)
}

// Refresh godoc
// @Summary Read flat rate shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Success 200 {object} responses.ResponseShippingFlatRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/flat-rate/{id} [get]
func (h *HandlersShippingOptions) ReadShippingFlatRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingFlatRates{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadFlatRateByID(&modelMethod, &modelRates, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.FlatRate {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingFlatRate(c, http.StatusOK, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Read flat rate shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Success 200 {object} responses.ResponseShippingFlatRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/table-rate/{id} [get]
func (h *HandlersShippingOptions) ReadShippingTableRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingTableRates{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadTableRateByID(&modelMethod, &modelRates, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.TableRate {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingTableRate(c, http.StatusOK, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Update local pickup method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Param params body requests.RequestShippingLocalPickup true "Method Info"
// @Success 200 {object} responses.ResponseShippingLocalPickup
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
	if err := methRepo.ReadByID(&modelMethod, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.PickUp {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := methService.UpdateShippingLocalPickup(req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingLocalPickup(c, http.StatusOK, modelMethod)
}

// Refresh godoc
// @Summary Update free shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Param params body requests.RequestShippingFree true "Method Info"
// @Success 200 {object} responses.ResponseShippingFree
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/free/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingFree(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingFree)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadByID(&modelMethod, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.FreeShipping {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not free shipping method.")
	}
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := methService.UpdateShippingFree(req, &modelMethod); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	return responses.NewResponseShippingFree(c, http.StatusOK, modelMethod)
}

// Refresh godoc
// @Summary Update table rate shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Param params body requests.RequestTableRate true "Class Info"
// @Success 201 {object} responses.ResponseShippingTableRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/rate/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestTableRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelRate := models.ShippingTableRates{}
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	if err := tableService.Update(methodID, req, &modelRate); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to update shipping rate.")
	}
	return responses.NewResponseTableRate(c, http.StatusOK, modelRate)
}

// Refresh godoc
// @Summary Update flat rate shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Param params body requests.RequestShippingFlatRate true "Method Info"
// @Success 200 {object} responses.ResponseShippingFlatRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/flat-rate/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingFlatRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingFlatRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingFlatRates{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadFlatRateByID(&modelMethod, &modelRates, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.FlatRate {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := methService.UpdateShippingFlatRate(req, &modelMethod, &modelRates); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingFlatRate(c, http.StatusOK, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Update flat rate shipping method
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Method ID"
// @Param params body requests.RequestShippingFlatRate true "Method Info"
// @Success 200 {object} responses.ResponseShippingFlatRate
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/table-rate/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingTableRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestShippingTableRate)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelMethod := models.ShippingMethods{}
	modelRates := []models.ShippingTableRates{}
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	if err := methRepo.ReadTableRateByID(&modelMethod, &modelRates, methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No method has this id.")
	}
	if modelMethod.Method != utils.TableRate {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This method is not local pickup method.")
	}
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	if err := methService.UpdateShippingTableRate(req, &modelMethod, &modelRates); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Shipping zone in this method doesn't exist.")
	}
	classIDs := []uint64{}
	for _, modelRate := range modelRates {
		classIDs = append(classIDs, modelRate.ClassID)
	}
	modelClasses := []models.ShippingClasses{}
	classRepo := repositories.NewRepositoryShippingClass(h.server.DB)
	classRepo.ReadByIDs(&modelClasses, classIDs)
	return responses.NewResponseShippingTableRate(c, http.StatusOK, modelMethod, modelRates, modelClasses)
}

// Refresh godoc
// @Summary Update shipping class
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Class ID"
// @Param params body requests.RequestShippingClass true "Class Info"
// @Success 200 {object} responses.ResponseShippingClass
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

// Refresh godoc
// @Summary Update shipping zone
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Zone ID"
// @Param params body requests.RequestShippingZone true "Zone Info"
// @Success 200 {object} responses.ResponseShippingZone
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
// @Summary Delete table rate shipping method to store
// @Tags Shipping Options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Rate ID"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/shipping/rate/{id} [delete]
func (h *HandlersShippingOptions) DeleteShippingRate(c echo.Context) error {
	methodID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	if err := tableService.Delete(methodID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to delete rate.")
	}
	return responses.MessageResponse(c, http.StatusOK, "Shipping rate successfullly deleted.")
}
