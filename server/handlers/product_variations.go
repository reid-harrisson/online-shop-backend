package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodvarsvc "OnlineStoreBackend/services/product_variations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersProductVariations struct {
	server *s.Server
}

func NewHandlersProductVariations(server *s.Server) *HandlersProductVariations {
	return &HandlersProductVariations{server: server}
}

// Refresh godoc
// @Summary Create product variation
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param params body requests.RequestProductVariation true "Variation Info"
// @Success 201 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation [post]
func (h *HandlersProductVariations) Create(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	req := new(requests.RequestProductVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Stock level and price are required.")
	}

	modelVar := models.ProductVariations{}
	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.Create(&modelVar, req, productID)

	return responses.NewResponseProductVariation(c, http.StatusCreated, modelVar)
}

// Refresh godoc
// @Summary Read all product variation in store
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} responses.ResponseProductVariationsInStore
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation/store [get]
func (h *HandlersProductVariations) ReadVariationsInStore(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelVars := make([]models.ProductVariationsInStore, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByStore(&modelVars, storeID)

	return responses.NewResponseProductVariationsInStore(c, http.StatusOK, modelVars)
}

// Refresh godoc
// @Summary Read all product variation in product
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} []responses.ResponseProductVariationsInProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation/product [get]
func (h *HandlersProductVariations) ReadVariationsInProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelVars := make([]models.ProductVariationsInProduct, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByProduct(&modelVars, productID)

	return responses.NewResponseProductVariationsInProduct(c, http.StatusOK, modelVars)
}

// Refresh godoc
// @Summary Update product variation stock level
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param stock_level query string true "Stock Level"
// @Success 200 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation/stock-level/{id} [put]
func (h *HandlersProductVariations) UpdateStockLevel(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	stockLevel, _ := strconv.ParseFloat(c.QueryParam("stock_level"), 64)

	modelVar := models.ProductVariations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.UpdateStockLevel(&modelVar, stockLevel)

	return responses.NewResponseProductVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Update product variation
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param params body requests.RequestProductVariation true "Variation Info"
// @Success 200 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation/{id} [put]
func (h *HandlersProductVariations) Update(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestProductVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Stock level and price are required.")
	}

	modelVar := models.ProductVariations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.Update(&modelVar, req)

	return responses.NewResponseProductVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Delete product variation by ID
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation/{id} [delete]
func (h *HandlersProductVariations) Delete(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelVar := models.ProductVariations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.Delete(variationID)

	return responses.MessageResponse(c, http.StatusOK, "Succesfully deleted")
}
