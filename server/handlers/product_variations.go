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
// @Summary Read all product variation
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation [get]
func (h *HandlersProductVariations) ReadAll(c echo.Context) error {
	modelVars := make([]models.ProductVariations, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadAllVariations(&modelVars)

	return responses.NewResponseProductVariationWithProduct(c, http.StatusOK, modelVars)
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
	varRepo.ReadVariationByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.UpdateStockLevel(&modelVar, stockLevel)

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
	varRepo.ReadVariationByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceProductVariation(h.server.DB)
	varService.Delete(variationID)

	return responses.MessageResponse(c, http.StatusOK, "Succesfully deleted")
}
