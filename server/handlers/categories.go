package handlers

import (
	"OnlineStoreBackend/models"
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
// @Router /store/api/v1/category [post]
func (h *HandlersCategories) CreateCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategory := models.Categories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByName(&modelCategory, req.Name, storeID)
	if modelCategory.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category already exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Create(&modelCategory, req, storeID)

	modelCategories := make([]models.CategoriesWithChildren, 0)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseCategories(c, http.StatusCreated, modelCategories)
}

// Refresh godoc
// @Summary Read category
// @Tags Category Actions
// @Accept json
// @Produce json
// @Param store_id path int true "Store ID"
// @Success 200 {object} []responses.ResponseCategoryWithChildren
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/category [get]
func (h *HandlersCategories) ReadCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCategories := make([]models.CategoriesWithChildren, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
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
// @Router /store/api/v1/category/{id} [put]
func (h *HandlersCategories) UpdateCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	req := new(requests.RequestCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategory := models.Categories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	if err := cateRepo.ReadByID(&modelCategory, categoryID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category doesn't exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Update(&modelCategory, req)

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
// @Router /store/api/v1/category/{id} [delete]
func (h *HandlersCategories) DeleteCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelCategory := models.Categories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	if err := cateRepo.ReadByID(&modelCategory, categoryID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category doesn't exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Delete(categoryID)

	return responses.MessageResponse(c, http.StatusOK, "Category successfully deleted")
}
