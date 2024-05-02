package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersTaxs struct {
	server *s.Server
}

func NewHandlersTaxs(server *s.Server) *HandlersTaxs {
	return &HandlersTaxs{server: server}
}

// Refresh godoc
// @Summary View tax rates
// @Tags Tax Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseTaxSetting
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/tax [get]
func (h *HandlersTaxs) ReadTaxSetting(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	taxRepo := repositories.NewRepositoryTax(h.server.DB)
	modelTax := models.Taxes{}
	err = taxRepo.ReadByCustomerID(&modelTax, customerID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseTax(c, http.StatusOK, modelTax)
}
