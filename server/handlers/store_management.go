package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	catesvc "OnlineStoreBackend/services/categories"
	storesvc "OnlineStoreBackend/services/stores"
	tagsvc "OnlineStoreBackend/services/tags"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersStoreManagement struct {
	server *s.Server
}

func NewHandlersStoreManagement(server *s.Server) *HandlersStoreManagement {
	return &HandlersStoreManagement{server: server}
}

// Refresh godoc
// @Summary Create store
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestStore true "Store Info"
// @Success 201 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store [post]
func (h *HandlersStoreManagement) Create(c echo.Context) error {
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelStore := models.Stores{}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Create(&modelStore, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStore(c, http.StatusCreated, modelStore)
}

// Refresh godoc
// @Summary Add category
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Param params body requests.RequestCategory true "Category"
// @Success 201 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/category [post]
func (h *HandlersStoreManagement) CreateCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategory := models.StoreCategories{}
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
// @Summary Add tag
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Param params body requests.RequestTag true "Tag"
// @Success 201 {object} []responses.ResponseStoreTag
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/tag [post]
func (h *HandlersStoreManagement) CreateTag(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestTag)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelTag := models.StoreTags{}
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByName(&modelTag, req.Name, storeID)
	if modelTag.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This tag already exist in the store.")
	}
	tagService := tagsvc.NewServiceTag(h.server.DB)
	tagService.Create(req.Name, &modelTag, storeID)

	modelTags := make([]models.StoreTags, 0)
	tagRepo.ReadByStoreID(&modelTags, storeID)
	return responses.NewResponseStoreTags(c, http.StatusCreated, modelTags)
}

// Refresh godoc
// @Summary Read store
// @Tags Store Management
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store [get]
func (h *HandlersStoreManagement) Read(c echo.Context) error {
	modelStores := make([]models.Stores, 0)
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	if err := storeRepo.ReadAll(&modelStores); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	return responses.NewResponseStores(c, http.StatusOK, modelStores)
}

// Refresh godoc
// @Summary Read category
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/category [get]
func (h *HandlersStoreManagement) ReadCategory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseStoreCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Read tag
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} []responses.ResponseStoreTag
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/tag [post]
func (h *HandlersStoreManagement) ReadTag(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelTags := make([]models.StoreTags, 0)
	tagRepo := repositories.NewRepositoryTag(h.server.DB)
	tagRepo.ReadByStoreID(&modelTags, storeID)
	return responses.NewResponseStoreTags(c, http.StatusOK, modelTags)
}

// Refresh godoc
// @Summary Update store
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Param params body requests.RequestStore true "Store Info"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id} [put]
func (h *HandlersStoreManagement) Update(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestStore)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelStore := models.Stores{}
	storeRepo := repositories.NewRepositoryStore(h.server.DB)
	if err := storeRepo.ReadByID(&modelStore, storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Update(&modelStore, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseStore(c, http.StatusOK, modelStore)
}

// Refresh godoc
// @Summary Update category
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category_id path int true "Category ID"
// @Param id path int true "Store ID"
// @Param params body requests.RequestCategory true "Category"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/category/{category_id} [put]
func (h *HandlersStoreManagement) UpdateCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("category_id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestCategory)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCategory := models.StoreCategories{}
	cateRepo := repositories.NewRepositoryCategory(h.server.DB)
	cateRepo.ReadByCategoryID(&modelCategory, categoryID)
	if modelCategory.ID == 0 || modelCategory.StoreID != storeID {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This category doesn't exist in the store.")
	}
	cateService := catesvc.NewServiceCategory(h.server.DB)
	cateService.Update(&modelCategory, req)

	modelCategories := make([]models.StoreCategoriesWithChildren, 0)
	cateRepo.ReadByStoreID(&modelCategories, storeID)
	return responses.NewResponseStoreCategories(c, http.StatusOK, modelCategories)
}

// Refresh godoc
// @Summary Delete category
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param category_id path int true "Category ID"
// @Param id path int true "Store ID"
// @Success 200 {object} []responses.ResponseStoreCategory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id}/category/{category_id} [delete]
func (h *HandlersStoreManagement) DeleteCategory(c echo.Context) error {
	categoryID, _ := strconv.ParseUint(c.Param("category_id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelCategory := models.StoreCategories{}
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

// Refresh godoc
// @Summary Delete store
// @Tags Store Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/{id} [delete]
func (h *HandlersStoreManagement) Delete(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	storeService := storesvc.NewServiceStore(h.server.DB)
	if err := storeService.Delete(storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No store exist at this ID.")
	}
	return responses.ErrorResponse(c, http.StatusOK, "Store successfully deleted.")
}
