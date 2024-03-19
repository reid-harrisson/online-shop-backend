package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	prodsvc "OnlineStoreBackend/services/products"
	prodvarsvc "OnlineStoreBackend/services/variations"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersInventoryManagement struct {
	server *s.Server
}

func NewHandlersInventoryManagement(server *s.Server) *HandlersInventoryManagement {
	return &HandlersInventoryManagement{server: server}
}

// Refresh godoc
// @Summary Read inventory
// @Tags Inventory Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Success 200 {object} responses.ResponseInventory
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/inventory/{id} [get]
func (h *HandlersInventoryManagement) ReadInventory(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelInventories := []models.Inventories{}
	invenRepo := repositories.NewRepositoryInventory(h.server.DB)
	invenRepo.ReadInventories(&modelInventories, storeID)

	return responses.NewResponseInventory(c, http.StatusBadRequest, modelInventories)
}

// Refresh godoc
// @Summary Set minimum stock level of product
// @Tags Inventory Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Param minimum_stock_level query string true "Minimum Stock Level"
// @Success 200 {object} responses.ResponseProduct
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/inventory/min-stock-level/{id} [put]
func (h *HandlersInventoryManagement) UpdateMinimumStockLevel(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	minimumStockLevel, _ := strconv.ParseFloat(c.QueryParam("minimum_stock_level"), 64)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)

	prodService := prodsvc.NewServiceProduct(h.server.DB)
	if err := prodService.UpdateMinimumStockLevel(productID, minimumStockLevel); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	modelProduct.MinimumStockLevel = minimumStockLevel
	return responses.NewResponseProduct(c, http.StatusOK, modelProduct)
}

// Refresh godoc
// @Summary Set stock level of variation
// @Tags Inventory Management
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Variation ID"
// @Param stock_level query string true "Stock Level"
// @Success 200 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/inventory/stock-level/{id} [put]
func (h *HandlersInventoryManagement) UpdateStockLevel(c echo.Context) error {
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
