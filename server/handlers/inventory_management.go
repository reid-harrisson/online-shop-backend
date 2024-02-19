package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
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
// @Summary Create store
// @Tags inventory manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestStore true "Store Info"
// @Success 201 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /api/v1/store [post]
func (h *HandlersInventoryManagement) Create(c echo.Context) error {
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)

	if err := storeService.Create(&modelStore, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStore(c, http.StatusCreated, modelStore)
}

// Refresh godoc
// @Summary Enable or disable back order
// @Tags inventory manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /api/v1/store/backorder/{id} [put]
func (h *HandlersInventoryManagement) UpdateBackOrder(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.UpdateBackOrder(id, &modelStore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseBackOrder(c, http.StatusOK, modelStore.BackOrder)
}

// Refresh godoc
// @Summary Enable or disable stock tracking
// @Tags inventory manangement
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /api/v1/store/tracking/{id} [put]
func (h *HandlersInventoryManagement) UpdateStockTracking(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.UpdateStockTracking(id, &modelStore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStockTracking(c, http.StatusOK, modelStore.ShowOutOfStockProducts)
}
