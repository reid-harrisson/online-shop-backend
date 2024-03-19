package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
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

// Refresh godoc
// @Summary Read all combos
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo [get]
func (h *HandlersCombos) ReadAll(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCombos := []models.Combos{}
	modelItems := []models.ComboItems{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	combRepo.ReadByStoreID(&modelCombos, &modelItems, storeID)
	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Read approved combos
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/approve [get]
func (h *HandlersCombos) ReadApproved(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCombos := []models.Combos{}
	modelItems := []models.ComboItems{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	combRepo.ReadApproved(&modelCombos, &modelItems, storeID)
	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Read published combos
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/publish [get]
func (h *HandlersCombos) ReadPublished(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCombos := []models.Combos{}
	modelItems := []models.ComboItems{}
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	combRepo.ReadPublished(&modelCombos, &modelItems, storeID)
	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Update combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Combo ID"
// @Param params body requests.RequestCombo true "Combo"
// @Success 200 {object} responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/{id} [put]
func (h *HandlersCombos) Update(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	comboID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestCombo)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	status := utils.Draft
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	if err := combRepo.ReadStatus(&status, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}
	if status == utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo is in pending status.")
	}

	modelCombo := models.Combos{}
	modelItems := []models.ComboItems{}
	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.Update(&modelCombo, &modelItems, req, storeID, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}

	return responses.NewResponseCombo(c, http.StatusOK, modelCombo, modelItems)
}

// Refresh godoc
// @Summary Approve combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/approve/{id} [put]
func (h *HandlersCombos) UpdateApprove(c echo.Context) error {
	comboID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	status := utils.Draft
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	if err := combRepo.ReadStatus(&status, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}
	if status != utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo isn't published.")
	}

	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.UpdateStatus(utils.Approved, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "This combo is successfully approved.")
}

// Refresh godoc
// @Summary Reject combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/reject/{id} [put]
func (h *HandlersCombos) UpdateReject(c echo.Context) error {
	comboID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	status := utils.Draft
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	if err := combRepo.ReadStatus(&status, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}
	if status != utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo isn't published.")
	}

	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.UpdateStatus(utils.Rejected, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "This combo is rejected.")
}

// Refresh godoc
// @Summary Publish combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/publish/{id} [put]
func (h *HandlersCombos) UpdatePublish(c echo.Context) error {
	comboID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	status := utils.Draft
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	if err := combRepo.ReadStatus(&status, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}
	if status != utils.Draft {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo isn't changed.")
	}

	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.UpdateStatus(utils.Pending, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusOK, "This combo is successfully published.")
}

// Refresh godoc
// @Summary Delete combo
// @Tags Combo
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/combo/{id} [delete]
func (h *HandlersCombos) Delete(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	comboID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	combService := combsvc.NewServiceCombo(h.server.DB)
	if err := combService.Delete(storeID, comboID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This combo doesn't exist.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Combo succesfully deleted.")
}
