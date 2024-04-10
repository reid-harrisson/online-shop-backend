package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	storesvc "OnlineStoreBackend/services/stores"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersStores struct {
	server *s.Server
}

func NewHandlersStores(server *s.Server) *HandlersStores {
	return &HandlersStores{server: server}
}

// Refresh godoc
// @Summary Create store
// @Tags Store Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestStore true "Store Info"
// @Success 201 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store [post]
func (h *HandlersStores) Create(c echo.Context) error {
	userID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Create(&modelStore, req, userID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStore(c, http.StatusCreated, modelStore)
}

// Refresh godoc
// @Summary Read all stores
// @Tags Store Actions
// @Accept json
// @Produce json
// @Success 200 {object} []responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/all [get]
func (h *HandlersStores) ReadAll(c echo.Context) error {
	modelStores := make([]models.Stores, 0)
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	if err := storeRepo.ReadAll(&modelStores); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	return responses.NewResponseStores(c, http.StatusOK, modelStores)
}

// Refresh godoc
// @Summary Read stores by user
// @Tags Store Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/user [get]
func (h *HandlersStores) ReadByUser(c echo.Context) error {
	userID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)

	modelStores := make([]models.Stores, 0)
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	if err := storeRepo.ReadByUser(&modelStores, userID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	return responses.NewResponseStores(c, http.StatusOK, modelStores)
}

// Refresh godoc
// @Summary Update store
// @Tags Store Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Param params body requests.RequestStore true "Store Info"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id} [put]
func (h *HandlersStores) Update(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelStore := models.Stores{}
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	if err := storeRepo.ReadByID(&modelStore, storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Update(&modelStore, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStore(c, http.StatusOK, modelStore)
}

// Refresh godoc
// @Summary Delete store
// @Tags Store Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id} [delete]
func (h *HandlersStores) Delete(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Delete(storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	return responses.ErrorResponse(c, http.StatusOK, "Store successfully deleted.")
}
