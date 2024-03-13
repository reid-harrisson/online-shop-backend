package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	combsvc "OnlineStoreBackend/services/combos"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCombos struct {
	server *s.Server
}

func NewHandlersCombos(server *s.Server) *HandlersCombos {
	return &HandlersCombos{server: server}
}

// Refresh godoc
// @Summary Create combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCombo true "Combo"
// @Success 201 {object} responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo [post]
func (h *HandlersCombos) Create(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestCombo)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCombo := models.Combos{}
	modelItems := []models.ComboItems{}
	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.Create(&modelCombo, &modelItems, req, storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This store doesn't exist.")
	}

	return responses.NewResponseCombo(c, http.StatusCreated, modelCombo, modelItems)
}
