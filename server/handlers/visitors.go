package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	vistsvc "OnlineStoreBackend/services/visitors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlersVisitors struct {
	server *s.Server
}

func NewHandlersVisitors(server *s.Server) *HandlersVisitors {
	return &HandlersVisitors{server: server}
}

// Refresh godoc
// @Summary Create visitor
// @Tags Visitor
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestVisitor true "Vistor"
// @Success 201 {object} responses.ResponseVisitor
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/visit [post]
func (h *HandlersVisitors) Create(c echo.Context) error {
	req := new(requests.RequestVisitor)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelVisitor := models.Visitors{}
	vistService := vistsvc.NewServiceVisitor(h.server.DB)
	vistService.Create(&modelVisitor, req)
	return responses.NewResponseVisitor(c, http.StatusCreated, modelVisitor)
}
