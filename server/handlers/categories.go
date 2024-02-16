package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	catesvc "OnlineStoreBackend/services/categories"
	"fmt"
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
// @Summary Add Category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCategory true "Category"
// @Success 201 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /api/v1/category [post]
func (h *HandlersCategories) CreateCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategory := models.BaseCategories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByName(&modelCategory, req.Name, storeID)
	if modelCategory.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category already exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Create(&modelCategory, req, storeID)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseStoreCategories(c, http.StatusCreated, modelCategories)
}

// Refresh godoc
// @Summary Read Category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /api/v1/category [get]
func (h *HandlersCategories) ReadCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	fmt.Println(modelCategories)
	return responses.NewResponseStoreCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Edit Category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param store_id query int true "Store ID"
// @Param parent_id query int false "Parent ID"
// @Param name query string false "Name"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /api/v1/category/{id} [put]
func (h *HandlersCategories) UpdateCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	name := c.QueryParam("name")
	parentID, _ := strconv.ParseUint(c.QueryParam("parent_id"), 10, 64)

	modelCategory := models.BaseCategories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByCategoryID(&modelCategory, categoryID)
	if modelCategory.ID == 0 || modelCategory.StoreID != storeID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category doesn't exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Update(&modelCategory, name, parentID)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseStoreCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Delete Category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /api/v1/category/{id} [delete]
func (h *HandlersCategories) DeleteCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCategory := models.BaseCategories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByCategoryID(&modelCategory, categoryID)
	if modelCategory.ID == 0 || modelCategory.StoreID != storeID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category doesn't exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Delete(categoryID)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseStoreCategories(c, http.StatusOK, modelCategories)
}
