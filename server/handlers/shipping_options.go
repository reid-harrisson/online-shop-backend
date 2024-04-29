package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	eh "OnlineStoreBackend/pkgs/error"
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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/shipping/rate [post]
func (h *HandlersShippingOptions) CreateShippingRate(c echo.Context) error {
	req := new(requests.RequestTableRate)

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelMethod := models.ShippingMethods{}
	modelRate := models.ShippingTableRates{}

	// Create shipping table method
	methService := methsvc.NewServiceShippingMethod(h.server.DB)
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	err = methService.Create(storeID, &modelMethod)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Create shipping table rate
	err = tableService.Create(uint64(modelMethod.ID), req, &modelRate)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/shipping/rate [get]
func (h *HandlersShippingOptions) ReadShippingRate(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelRates := []models.ShippingTableRates{}

	// Read shipping table rate
	methRepo := repositories.NewRepositoryShippingMethod(h.server.DB)
	err = methRepo.ReadRates(&modelRates, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/shipping/rate/{id} [put]
func (h *HandlersShippingOptions) UpdateShippingRate(c echo.Context) error {
	req := new(requests.RequestTableRate)

	methodID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelRate := models.ShippingTableRates{}

	// Update shipping table rate
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	err = tableService.Update(methodID, req, &modelRate)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/shipping/rate/{id} [delete]
func (h *HandlersShippingOptions) DeleteShippingRate(c echo.Context) error {
	methodID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Delete shipping table rate
	tableService := tablesvc.NewServiceShippingTableRate(h.server.DB)
	err = tableService.Delete(methodID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.ShippingRateDeleted)
}
