package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	tagsvc "OnlineStoreBackend/services/tags"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersTags struct {
	server *s.Server
}

func NewHandlersTags(server *s.Server) *HandlersTags {
	return &HandlersTags{server: server}
}

// Refresh godoc
// @Summary Add tag
// @Tags Tag Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestTag true "Tag"
// @Success 201 {object} []responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/tag [post]
func (h *HandlersTags) CreateTag(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestTag)
	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	if err := tagRepo.ReadByNameAndStoreID(&modelTag, req.Name, storeID); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.TagDuplicated)
	}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	err = tagService.Create(&modelTag, req.Name, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelTags := make([]models.Tags, 0)
	err = tagRepo.ReadByStoreID(&modelTags, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseTags(c, http.StatusCreated, modelTags)
}

// Refresh godoc
// @Summary Read tag
// @Tags Tag Actions
// @Accept json
// @Produce json
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/tag [get]
func (h *HandlersTags) ReadTag(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTags := make([]models.Tags, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	err = tagRepo.ReadByStoreID(&modelTags, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseTags(c, http.StatusOK, modelTags)
}

// Refresh godoc
// @Summary Update tag
// @Tags Tag Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Tag ID"
// @Param params body requests.RequestTag true "Tag"
// @Success 200 {object} responses.ResponseTag
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/tag/{id} [put]
func (h *HandlersTags) UpdateTag(c echo.Context) error {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestTag)
	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelTag := models.Tags{}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	err = tagService.Update(tagID, &modelTag, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseTag(c, http.StatusOK, modelTag)
}

// Refresh godoc
// @Summary Delete tag
// @Tags Tag Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param id path int true "Tag ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/tag/{id} [delete]
func (h *HandlersTags) DeleteTag(c echo.Context) error {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	err = tagService.Delete(tagID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, "Tag succesfully deleted")
}
