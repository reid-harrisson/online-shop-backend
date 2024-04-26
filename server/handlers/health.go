package handlers

import (
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	server *s.Server
}

func NewHealthHandler(server *s.Server) *HealthHandler {
	return &HealthHandler{server: server}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Check server's health
// @ID health-check
// @Tags Health Check Actions
// @Accept json
// @Produce json
// @Success 200 {object} responses.Data
// @Router /store/api/v1/health [get]
func (healthHandler *HealthHandler) HealthCheck(c echo.Context) error {
	return responses.MessageResponse(c, http.StatusOK, "Server is running!")
}
