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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/store [post]
func (h *HandlersStores) Create(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestStore)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if !utils.ValidateEmailAddress(req.ContactEmail) {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidEmailAddress)
	}

	if !utils.ValidatePhoneNumber(req.ContactPhone) {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidPhoneNumber)
	}

	// Create store with user ID
	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	err = storeService.Create(&modelStore, req, userID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/store/all [get]
func (h *HandlersStores) ReadAll(c echo.Context) error {
	modelStores := make([]models.Stores, 0)

	// Read all stores
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	err := storeRepo.ReadAll(&modelStores)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/store/user [get]
func (h *HandlersStores) ReadByUser(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read store by user ID
	modelStores := make([]models.Stores, 0)
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	err = storeRepo.ReadByUser(&modelStores, userID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/store/{id} [put]
func (h *HandlersStores) Update(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// Update store by ID
	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	err = storeService.Update(&modelStore, req, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
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
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/store/{id} [delete]
func (h *HandlersStores) Delete(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Delete store by ID
	storeService := storesvc.NewServiceStore(h.server.DB)
	err = storeService.Delete(storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.ErrorResponse(c, http.StatusOK, constants.SuccessDeleteStore)
}
