package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cartsvc "OnlineStoreBackend/services/cart_items"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersShoppingCart struct {
	server *s.Server
}

func NewHandlersCart(server *s.Server) *HandlersShoppingCart {
	return &HandlersShoppingCart{server: server}
}

// Refresh godoc
// @Summary Add cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCartItem true "Variation Info"
// @Success 201 {object} responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [post]
func (h *HandlersShoppingCart) Create(c echo.Context) error {
	req := new(requests.RequestCartItem)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	if err := prodRepo.ReadByID(&modelProduct, req.ProductID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product doesn't exist at this ID.")
	}
	if modelProduct.Status != utils.Approved {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This product isn't approved.")
	}

	modelVar := models.ProductVariations{}
	modelVals := []models.ProductAttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryProductAttributeValue(h.server.DB)
	valRepo.ReadByIDs(&modelVals, req.ValueIDs)
	mapVal := map[uint64]string{}
	for _, modelVal := range modelVals {
		if mapVal[modelVal.AttributeID] == "" {
			mapVal[modelVal.AttributeID] = modelVal.AttributeValue
		} else {
			return responses.ErrorResponse(c, http.StatusBadRequest, "Attribute value's duplicated.")
		}
	}
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	varRepo.ReadByAttributeValueIDs(&modelVar, req.ValueIDs, req.ProductID)

	if modelVar.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This variation doesn't exist in product.")
	}

	variationID := uint64(modelVar.ID)

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByInfo(&modelItem, variationID, req.CustomerID)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.Create(&modelItem, req.CustomerID, &modelVar, float64(req.Quantity))

	return responses.NewResponseCartItem(c, http.StatusCreated, modelItem)
}

// Refresh godoc
// @Summary Read count
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCartCount
// @Failure 400 {object} responses.Error
// /@Router /store/api/v1/cart/count [get]
func (h *HandlersShoppingCart) ReadCount(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)

	modelCount := models.CartCount{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadItemCount(&modelCount, customerID)
	return responses.NewResponseCartCount(c, http.StatusOK, modelCount)
}

// Refresh godoc
// @Summary Read cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [get]
func (h *HandlersShoppingCart) Read(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo.ReadDetail(&modelItems, customerID)
	return responses.NewResponseCart(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Edit cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Param quantity query string true "Quantity"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart/{id} [put]
func (h *HandlersShoppingCart) UpdateQuantity(c echo.Context) error {
	cartID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	quantity, _ := strconv.ParseFloat(c.QueryParam("quantity"), 64)

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByID(&modelItem, cartID)
	if modelItem.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No cart item exist at this ID.")
	}

	modelVar := models.ProductVariations{}
	prodRepo := repositories.NewRepositoryVariation(h.server.DB)
	prodRepo.ReadByID(&modelVar, modelItem.VariationID)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.UpdateQuantity(&modelItem, modelVar, quantity)

	return responses.NewResponseCartItem(c, http.StatusOK, modelItem)
}

// Refresh godoc
// @Summary Delete a item
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart/{id} [delete]
func (h *HandlersShoppingCart) Delete(c echo.Context) error {
	cartID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByID(&modelItem, cartID)

	if modelItem.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No cart item exists at this ID.")
	}

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.Delete(cartID)

	return responses.MessageResponse(c, http.StatusOK, "This cart item successfullly deleted.")
}

// Refresh godoc
// @Summary Delete all items in cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [delete]
func (h *HandlersShoppingCart) DeleteAll(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)

	modelItems := make([]models.CartItems, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByCustomerID(&modelItems, customerID)

	if len(modelItems) == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "There is no items in cart.")
	}

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.DeleteAll(customerID)
	return responses.MessageResponse(c, http.StatusOK, "All cart items successfully deleted.")
}
