package handlers

import (
	"OnlineStoreBackend/models"
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
// @Router /store/api/v1/tag [post]
func (h *HandlersTags) CreateTag(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByName(&modelTag, req.Name, storeID)
	if modelTag.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This tag already exist in the store.")
	}
	tagService := tagsvc.NewServiceTag(h.server.DB)
	tagService.Create(&modelTag, req.Name, storeID)

	modelTags := make([]models.Tags, 0)
	tagRepo.ReadByStoreID(&modelTags, storeID)
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
// @Router /store/api/v1/tag [get]
func (h *HandlersTags) ReadTag(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelTags := make([]models.Tags, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByStoreID(&modelTags, storeID)
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
// @Router /store/api/v1/tag/{id} [put]
func (h *HandlersTags) UpdateTag(c echo.Context) error {
	tagID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	if err := tagRepo.ReadByID(&modelTag, tagID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This tag doesn't exist.")
	}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	tagService.Update(&modelTag, req)

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
// @Router /store/api/v1/tag/{id} [delete]
func (h *HandlersTags) DeleteTag(c echo.Context) error {
	tagID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelTag := models.Tags{}
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	if err := tagRepo.ReadByID(&modelTag, tagID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This tag doesn't exist.")
	}

	tagService := tagsvc.NewServiceTag(h.server.DB)
	tagService.Delete(tagID)

	return responses.MessageResponse(c, http.StatusOK, "Tag succesfully deleted")
}
