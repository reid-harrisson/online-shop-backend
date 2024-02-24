package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	storesvc "OnlineStoreBackend/services/stores"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlersGeneralStoreOffering struct {
	server *s.Server
}

func NewHandlersGeneralStoreOffering(server *s.Server) *HandlersGeneralStoreOffering {
	return &HandlersGeneralStoreOffering{server: server}
}

// Refresh godoc
// @Summary Create store
// @Tags General Store Offering
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestStore true "Store Info"
// @Success 201 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store [post]
func (h *HandlersGeneralStoreOffering) Create(c echo.Context) error {
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
