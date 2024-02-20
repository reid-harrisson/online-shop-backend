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

func NewHandlersTaxSettings(server *s.Server) *HandlersTaxSettings {
	return &HandlersTaxSettings{server: server}
}

// Refresh godoc
// @Summary View tax rates
// @Tags tax settings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseTaxSetting
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/tax [get]
func (h *HandlersTaxSettings) ReadTaxSetting(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	modelTax := models.TaxSettings{}
	if err := taxRepo.ReadTaxSetting(&modelTax, customerID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseTaxSetting(c, http.StatusOK, modelTax)
}
