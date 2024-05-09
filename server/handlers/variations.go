package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	eh "OnlineStoreBackend/pkgs/error"
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
	req := new(requests.RequestVariation)

	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Check duplicate variation
	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByAttributeValueIDs(&modelVar, req.AttributeValueIDs, productID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 && statusCode != http.StatusNotFound {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedVariation)
	}

	// Create variation
	modelVar = models.Variations{}
	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.Create(&modelVar, req, productID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseVariation(c, http.StatusCreated, modelVar)
}

// Refresh godoc
// @Summary Read product variation by attribute values
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Param product_id query int true "Product ID"
// @Param attribute_value_ids query string false "Attribute Value IDs"
// @Success 200 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation [get]
func (h *HandlersVariations) ReadByAttributeValues(c echo.Context) error {
	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	values := strings.Split(c.QueryParam("attribute_value_ids"), ",")
	valueIDs := []uint64{}

	for _, value := range values {
		if value != "" {
			valueID, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
			}
			valueIDs = append(valueIDs, valueID)
		}
	}

	// Read variatio nby attribute value id
	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByAttributeValueIDs(&modelVar, valueIDs, productID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseVariation(c, http.StatusOK, modelVar)
}

// Refresh godoc
// @Summary Read all product variation in product
// @Tags Variation Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Success 200 {object} []responses.ResponseVariationsInProduct
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/variation/product [get]
func (h *HandlersVariations) ReadByProduct(c echo.Context) error {
	productID, err := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read variation by product
	modelVars := make([]models.VariationsWithAttributeValue, 0)
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByProduct(&modelVars, productID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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
	variationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelVar := models.Variations{}

	// Read variation by id
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByID(&modelVar, variationID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Update back-order
	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.UpdateBackOrder(&modelVar)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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
	req := new(requests.RequestVariation)

	variationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read variation by id
	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByID(&modelVar, variationID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Check duplicate variation
	err = varRepo.ReadByAttributeValueIDs(&models.Variations{}, req.AttributeValueIDs, modelVar.ProductID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 && statusCode != http.StatusNotFound {
		return responses.ErrorResponse(c, statusCode, message)
	} else if statusCode == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicatedVariation)
	}

	// Update variation
	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.Update(&modelVar, req)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

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
	variationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read variation by id
	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByID(&modelVar, variationID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Delete variation
	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.Delete(variationID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.VariationDeleted)
}
