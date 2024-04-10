package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersTaxSettings struct {
	server *s.Server
}

func NewHandlersTaxs(server *s.Server) *HandlersTaxSettings {
	return &HandlersTaxSettings{server: server}
}

// Refresh godoc
// @Summary View tax rates
// @Tags Tax Settings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseTaxSetting
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/tax [get]
func (h *HandlersTaxSettings) ReadTaxSetting(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)

	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	modelTax := models.Taxes{}
	if err := taxRepo.ReadByCustomerID(&modelTax, customerID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseTaxSetting(c, http.StatusOK, modelTax)
}
