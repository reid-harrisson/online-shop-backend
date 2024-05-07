package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	catesvc "OnlineStoreBackend/services/categories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCategories struct {
	server *s.Server
}

func NewHandlersCategories(server *s.Server) *HandlersCategories {
	return &HandlersCategories{server: server}
}

// Refresh godoc
// @Summary Add category
// @Tags Category Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCategory true "Category"
// @Success 201 {object} []responses.ResponseCategoryWithChildren
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/category [post]
func (h *HandlersCategories) CreateCategory(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCategory)
	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	cateRepo := repositories.NewRepositoryCategory(h.server.DB)

	// check match between input store_id and parent_id's store_id(same store)
	modelCategory := models.Categories{}
	if err := cateRepo.ReadByID(&modelCategory, req.ParentID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if modelCategory.StoreID != storeID {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// check duplicate
	if err := cateRepo.ReadByNameAndStoreIDAndParentID(&modelCategory, req.Name, storeID, req.ParentID); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.CategoryDuplicated)
	}

	modelCategoryCreate := models.Categories{}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	err = cateService.Create(&modelCategoryCreate, req, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	modelCategoriesWithChildren := make([]models.CategoriesWithChildren, 0)
	err = cateRepo.ReadByStoreID(&modelCategoriesWithChildren, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCategories(c, http.StatusCreated, modelCategoriesWithChildren)
}

// Refresh godoc
// @Summary Read category
// @Tags Category Actions
// @Accept json
// @Produce json
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseCategoryWithChildren
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/category [get]
func (h *HandlersCategories) ReadCategory(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCategories := make([]models.CategoriesWithChildren, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	err = cateRepo.ReadByStoreID(&modelCategories, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Update category
// @Tags Category Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param params body requests.RequestCategory true "Category"
// @Success 200 {object} responses.ResponseCategory
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/category/{id} [put]
func (h *HandlersCategories) UpdateCategory(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCategory)
	err = c.Bind(req)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCategory := models.Categories{}

	cateService := catesvc.NewServiceCategory(h.server.DB)
	err = cateService.Update(categoryID, &modelCategory, req)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCategory(c, http.StatusOK, modelCategory)
}

// Refresh godoc
// @Summary Delete category
// @Tags Category Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/category/{id} [delete]
func (h *HandlersCategories) DeleteCategory(c echo.Context) error {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCategory := models.Categories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	err = cateRepo.ReadByID(&modelCategory, categoryID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	cateService := catesvc.NewServiceCategory(h.server.DB)
	err = cateService.Delete(categoryID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, "Category successfully deleted")
}
