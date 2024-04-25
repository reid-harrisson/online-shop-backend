package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	errhandle "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodsvc "OnlineStoreBackend/services/products"
	stocksvc "OnlineStoreBackend/services/stock_tracks"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersInventory struct {
	server *s.Server
}

func NewHandlersInventory(server *s.Server) *HandlersInventory {
	return &HandlersInventory{server: server}
}

// Refresh godoc
// @Summary Read inventory
// @Tags Inventory Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseInventory
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/inventory/{id} [get]
func (h *HandlersInventory) ReadInventory(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelInventories := []models.Inventories{}
	invenRepo := repositories.NewRepositoryInventory(h.server.DB)
	err = invenRepo.ReadInventories(&modelInventories, storeID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseInventory(c, http.StatusBadRequest, modelInventories)
}

// Refresh godoc
// @Summary Get minimum stock level of product
// @Tags Inventory Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} responses.ResponseMinimumStockLevel
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/inventory/min-stock-level/{id} [get]
func (h *HandlersInventory) GetMinimumStockLevel(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	minimumStockLevel := float64(0)

	prodRepo := repositories.NewRepositoryProduct(h.server.DB)

	err = prodRepo.GetMinimumStockLevel(&minimumStockLevel, productID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseMinimumStockLevel(c, http.StatusOK, minimumStockLevel)
}

// Refresh godoc
// @Summary Set minimum stock level of product
// @Tags Inventory Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param minimum_stock_level query string true "Minimum Stock Level"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/inventory/min-stock-level/{id} [put]
func (h *HandlersInventory) SetMinimumStockLevel(c echo.Context) error {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	minimumStockLevel, err := strconv.ParseFloat(c.QueryParam("minimum_stock_level"), 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	err = prodService.UpdateMinimumStockLevel(productID, minimumStockLevel)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, "Successfully set minimum stock level!")
}

// Refresh godoc
// @Summary Get stock level of variation
// @Tags Inventory Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param stock_level query string true "Stock Level"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/inventory/stock-level/{id} [get]
func (h *HandlersInventory) GetStockLevel(c echo.Context) error {
	variationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByID(&modelVar, variationID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseStockLevel(c, http.StatusOK, modelVar.StockLevel)
}

// Refresh godoc
// @Summary Set stock level of variation
// @Tags Inventory Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param stock_level query string true "Stock Level"
// @Success 200 {object} responses.ResponseVariation
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/inventory/stock-level/{id} [put]
func (h *HandlersInventory) SetStockLevel(c echo.Context) error {
	variationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	stockLevel, err := strconv.ParseFloat(c.QueryParam("stock_level"), 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelVar := models.Variations{}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByID(&modelVar, variationID)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	prevStockLevel := modelVar.StockLevel

	varService := prodvarsvc.NewServiceVariation(h.server.DB)
	err = varService.UpdateStockLevel(&modelVar, stockLevel)
	if statusCode, message := errhandle.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Track Stock Level
	stockService := stocksvc.NewServiceStockTrack(h.server.DB)
	stockService.Create(models.StockTracks{
		ProductID:   modelVar.ProductID,
		VariationID: variationID,
		Change:      stockLevel - prevStockLevel,
		Event:       utils.ProductWarhousing,
	})

	return responses.NewResponseVariation(c, http.StatusOK, modelVar)
}
