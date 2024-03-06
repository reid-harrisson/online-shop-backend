package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	storesvc "OnlineStoreBackend/services/stores"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersInventoryManagement struct {
	server *s.Server
}

func NewHandlersInventoryManagement(server *s.Server) *HandlersInventoryManagement {
	return &HandlersInventoryManagement{server: server}
}

// Refresh godoc
// @Summary Show or hide out of stock
// @Tags Inventory Manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/inventory/out-of-stock/{id} [put]
func (h *HandlersInventoryManagement) UpdateShowOutOfStockStatus(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelStore := models.Stores{}

	repositoryInventory := repositories.NewRepositoryInventory(h.server.DB)
	if err := repositoryInventory.ReadOne(&modelStore, id); err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Data index not found")
	}

	storeService := storesvc.NewServiceStore(h.server.DB)

	if err := storeService.UpdateOutOfStockStatus(id, &modelStore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseOutOfStockStatus(c, http.StatusOK, modelStore.ShowOutOfStockStatus)
}

// Refresh godoc
// @Summary Show or hide stock level
// @Tags Inventory Manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/inventory/stock-level/{id} [put]
func (h *HandlersInventoryManagement) UpdateShowStockLevelStatus(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelStore := models.Stores{}

	repositoryInventory := repositories.NewRepositoryInventory(h.server.DB)
	if err := repositoryInventory.ReadOne(&modelStore, id); err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Data index not found")
	}

	storeService := storesvc.NewServiceStore(h.server.DB)

	if err := storeService.UpdateStockLevelStatus(id, &modelStore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStockLevelStatus(c, http.StatusOK, modelStore.ShowStockLevelStatus)
}
