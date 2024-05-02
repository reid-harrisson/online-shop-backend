package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	etsvc "OnlineStoreBackend/services/email_templates"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersTemplates struct {
	server *s.Server
}

func NewHandlersTemplates(server *s.Server) *HandlersTemplates {
	return &HandlersTemplates{server: server}
}

// Refresh godoc
// @Summary Create email template
// @Tags Email Template Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestEmailTemplate true "Email Template Data"
// @Success 200 {object} responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/template [post]
func (h *HandlersTemplates) CreateTemplate(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestEmailTemplate)

	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	err = req.Validate()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTemplate := models.EmailTemplates{}
	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	err = temService.Create(&modelTemplate, req, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseEmailTemplate(c, http.StatusCreated, &modelTemplate)
}

// Refresh godoc
// @Summary Read email templates
// @Tags Email Template Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/template [get]
func (h *HandlersTemplates) ReadTemplate(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTemplates := make([]models.EmailTemplates, 0)

	temRepo := repositories.NewRepositoryEmailTemplate(h.server.DB)
	err = temRepo.ReadByStoreID(&modelTemplates, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseEmailTemplates(c, http.StatusOK, modelTemplates)
}

// Refresh godoc
// @Summary Update email template
// @Tags Email Template Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Template ID"
// @Param params body requests.RequestEmailTemplate true "Email Template Data"
// @Success 200 {object} responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/template/{id} [put]
func (h *HandlersTemplates) UpdateTemplate(c echo.Context) error {
	templateID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestEmailTemplate)

	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	err = req.Validate()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTemplate := models.EmailTemplates{}

	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	err = temService.Update(templateID, &modelTemplate, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseEmailTemplate(c, http.StatusOK, &modelTemplate)
}

// Refresh godoc
// @Summary Delete email template
// @Tags Email Template Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Template ID"
// @Success 200 {object} []responses.ResponseEmailTemplate
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/template/{id} [delete]
func (h *HandlersTemplates) DeleteTemplate(c echo.Context) error {
	templateID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	err = temService.Delete(templateID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, "Successfully deleted")
}
