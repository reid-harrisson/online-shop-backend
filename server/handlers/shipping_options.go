package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	methsvc "OnlineStoreBackend/services/shipping_methods"
	tablesvc "OnlineStoreBackend/services/shipping_table_rates"
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
// @Summary Add table rate shipping method to store
// @Tags Shipping Option Actions
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

	modelMethod := models.ShippingMethods{}
	modelRate := models.ShippingTableRates{}
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	methService.Create(storeID, &modelMethod)
	if err := tableService.Create(uint64(modelMethod.ID), req, &modelRate); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This condition already exist.")
	}
	return responses.NewResponseTableRate(c, http.StatusCreated, modelRate)
}

// Refresh godoc
// @Summary Read rates
// @Tags Shipping Option Actions
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
// @Summary Update table rate shipping method to store
// @Tags Shipping Option Actions
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
// @Summary Delete table rate shipping method to store
// @Tags Shipping Option Actions
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
