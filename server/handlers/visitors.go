package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
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
// @Tags Visitor Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestVisitor true "Vistor"
// @Success 201 {object} responses.ResponseVisitor
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/visit [post]
func (h *HandlersVisitors) Create(c echo.Context) error {
	req := new(requests.RequestVisitor)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelVisitor := models.Visitors{}
	vistService := vistsvc.NewServiceVisitor(h.server.DB)
	err := vistService.Create(&modelVisitor, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}
	return responses.NewResponseVisitor(c, http.StatusCreated, modelVisitor)
}
