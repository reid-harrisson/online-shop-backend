package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	eh "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	combsvc "OnlineStoreBackend/services/combos"
	"encoding/json"
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
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCombo true "Combo"
// @Success 201 {object} responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo [post]
func (h *HandlersCombos) Create(c echo.Context) error {
	req := new(requests.RequestCombo)

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var modelCombo = models.Combos{}
	var modelItems = make([]models.ComboItems, 0)

	// Initialize combo
	imageUrls, _ := json.Marshal(req.ImageUrls)

	modelCombo.StoreID = storeID
	modelCombo.DiscountAmount = req.DiscountAmount
	modelCombo.DiscountType = utils.DiscountTypeFromString(req.DiscountType)
	modelCombo.ImageUrls = string(imageUrls)
	modelCombo.Description = req.Description
	modelCombo.Title = req.Title
	modelCombo.Status = utils.Draft

	// Create combo
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.Create(&modelCombo, &modelItems, req, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCombo(c, http.StatusCreated, modelCombo, modelItems)
}

// Refresh godoc
// @Summary Read all combos
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo [get]
func (h *HandlersCombos) ReadAll(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelCombos = []models.Combos{}
	var modelItems = []models.ComboItems{}

	// Read all combos
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadByStoreID(&modelCombos, &modelItems, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Read approved combos
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/approved [get]
func (h *HandlersCombos) ReadApproved(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelCombos = []models.Combos{}
	var modelItems = []models.ComboItems{}

	// Read approved combo
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadApproved(&modelCombos, &modelItems, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Read published combos
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 201 {object} []responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/publish [get]
func (h *HandlersCombos) ReadPublished(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelCombos = []models.Combos{}
	var modelItems = []models.ComboItems{}

	// Read published combo
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadPublished(&modelCombos, &modelItems, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCombos(c, http.StatusOK, modelCombos, modelItems)
}

// Refresh godoc
// @Summary Update combo
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Combo ID"
// @Param params body requests.RequestCombo true "Combo"
// @Success 200 {object} responses.ResponseCombo
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/{id} [put]
func (h *HandlersCombos) Update(c echo.Context) error {
	req := new(requests.RequestCombo)

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	comboID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	status := utils.Draft

	// Read status
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadStatus(&status, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	if status == utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.ComboPending)
	}

	modelCombo := models.Combos{}
	var modelItems = []models.ComboItems{}

	// Update combo
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.Update(&modelCombo, &modelItems, req, storeID, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCombo(c, http.StatusOK, modelCombo, modelItems)
}

// Refresh godoc
// @Summary Approve combo
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/approve/{id} [put]
func (h *HandlersCombos) UpdateApprove(c echo.Context) error {
	comboID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := utils.Draft

	// Read status
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadStatus(&status, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	if status != utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.ComboNotPublished)
	}

	// Update status
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.UpdateStatus(utils.Approved, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.ComboApproved)
}

// Refresh godoc
// @Summary Reject combo
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/reject/{id} [put]
func (h *HandlersCombos) UpdateReject(c echo.Context) error {
	comboID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := utils.Draft

	// Read status
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadStatus(&status, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	if status != utils.Pending {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.ComboNotPublished)
	}

	// Update status
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.UpdateStatus(utils.Rejected, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.ComboRejected)
}

// Refresh godoc
// @Summary Publish combo
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/publish/{id} [put]
func (h *HandlersCombos) UpdatePublish(c echo.Context) error {
	comboID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	status := utils.Draft

	// Read status
	combRepo := repositories.NewRepositoryCombo(h.server.DB)
	err = combRepo.ReadStatus(&status, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	if status != utils.Draft {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.ComboNotChanged)
	}

	// Update status
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.UpdateStatus(utils.Pending, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.ComboPublished)
}

// Refresh godoc
// @Summary Delete combo
// @Tags Combo Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Combo ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/combo/{id} [delete]
func (h *HandlersCombos) Delete(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	comboID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Delete combo
	combService := combsvc.NewServiceCombo(h.server.DB)
	err = combService.Delete(storeID, comboID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.ComboDeleted)
}
