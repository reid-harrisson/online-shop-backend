package handlers

import (
	"OnlineStoreBackend/models"
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
// @Router /store/api/v1/store/template [post]
func (h *HandlersTemplates) CreateTemplate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestEmailTemplate)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTemplate := models.EmailTemplates{}
	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	if err := temService.Create(&modelTemplate, req, storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
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
// @Router /store/api/v1/store/template [get]
func (h *HandlersTemplates) ReadTemplate(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelTemplates := make([]models.EmailTemplates, 0)

	temRepo := repositories.NewRepositoryEmailTemplate(h.server.DB)
	temRepo.ReadByStoreID(&modelTemplates, storeID)

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
// @Router /store/api/v1/store/template/{id} [put]
func (h *HandlersTemplates) UpdateTemplate(c echo.Context) error {
	templateID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestEmailTemplate)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTemplate := models.EmailTemplates{}

	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	if err := temService.Update(templateID, &modelTemplate, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No template exist at this ID.")
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
// @Router /store/api/v1/store/template/{id} [delete]
func (h *HandlersTemplates) DeleteTemplate(c echo.Context) error {
	templateID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	temService := etsvc.NewServiceEmailTemplate(h.server.DB)
	if err := temService.Delete(templateID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No template exist at this ID.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Successfully deleted")
}
