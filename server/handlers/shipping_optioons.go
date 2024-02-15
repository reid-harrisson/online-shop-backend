package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	shipData "OnlineStoreBackend/services/shipping_data"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlersShippingOptions struct {
	server *s.Server
}

func NewHandlersShippingOptions(server *s.Server) *HandlersShippingOptions {
	return &HandlersShippingOptions{server: server}
}

// Refresh godoc
// @Summary Set shipping method
// @Tags shipping options
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestShippingMethod true "Shipping Method"
// @Success 200 {object} []responses.ResponseShippingData
// @Failure 400 {object} responses.Error
// @Router /api/v1/shipping-option/method [get]
func (h *HandlersShippingOptions) UpdateShippingMethod(c echo.Context) error {
	req := new(requests.RequestShippingMethod)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	shipDataService := shipData.CreateService(h.server.DB)
	modelShipData := models.ShippingData{}
	if err := shipDataService.Update(req, &modelShipData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseShippingData(c, http.StatusOK, modelShipData)
}
