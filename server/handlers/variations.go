package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type HandlersVariations struct {
	server *s.Server
}

func NewHandlersVariations(server *s.Server) *HandlersVariations {
	return &HandlersVariations{server: server}
}

// Refresh godoc
// @Summary Create product variation
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param params body requests.RequestVariation true "Variation Info"
// @Success 201 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation [post]
func (h *HandlersVariations) Create(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	req := new(requests.RequestVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Stock level and price are required.")
	}

	modelVar := models.Variations{}
	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	varService.Create(&modelVar, req, productID)

	return responses.NewResponseVariation(c, http.StatusCreated, modelVar)
}

// Refresh godoc
// @Summary Read product variation by attribute values
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param attribute_value_ids query string true "Attribute Value IDs"
// @Success 200 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation [get]
func (h *HandlersVariations) ReadByAttributeValues(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	values := strings.Split(c.QueryParam("attribute_value_ids"), ",")

	valueIDs := []uint64{}
	for _, value := range values {
		valueID, _ := strconv.ParseUint(value, 10, 64)
		valueIDs = append(valueIDs, valueID)
	}

	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByAttributeValueIDs(&modelVar, valueIDs, productID)

	return responses.NewResponseVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Read all product variation in product
// @Tags Variation Actions
// @Accept json
// @Produce json
// /@Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} []responses.ResponseVariationsInProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation/product [get]
func (h *HandlersVariations) ReadByProduct(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)

	modelVars := make([]models.VariationsWithAttributeValue, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByProduct(&modelVars, productID)

	return responses.NewResponseVariationsInProduct(c, http.StatusOK, modelVars)
}

// Refresh godoc
// @Summary Enable or disable back order
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Success 200 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation/back-order/{id} [put]
func (h *HandlersVariations) UpdateBackOrder(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelVar := models.Variations{}

	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	if err := varRepo.ReadByID(&modelVar, variationID); err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	varService.UpdateBackOrder(&modelVar)

	return responses.NewResponseVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Update product variation
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param params body requests.RequestVariation true "Variation Info"
// @Success 200 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation/{id} [put]
func (h *HandlersVariations) Update(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Stock level and price are required.")
	}

	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	varService.Update(&modelVar, req)

	return responses.NewResponseVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Delete product variation by ID
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation/{id} [delete]
func (h *HandlersVariations) Delete(c echo.Context) error {
	variationID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByID(&modelVar, variationID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "No record found")
	}

	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	varService.Delete(variationID)

	return responses.MessageResponse(c, http.StatusOK, "Succesfully deleted")
}
