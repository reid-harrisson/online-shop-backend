package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
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

func NewHandlersShoppingCart(server *s.Server) *HandlersShoppingCart {
	return &HandlersShoppingCart{server: server}
}

// Refresh godoc
// @Summary Add cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param product_id query int true "Product ID"
// @Param quantity query string true "Quantity"
// @Success 201 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [post]
func (h *HandlersShoppingCart) Create(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	quantity, _ := strconv.ParseFloat(c.QueryParam("quantity"), 64)

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, productID)
	if modelProduct.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No product exists at this product ID.")
	}

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByProductID(&modelItem, productID, customerID)
	if modelItem.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This cart item already exist in cart.")
	}
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.Create(&modelItem, customerID, modelProduct, quantity)

	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo.ReadDetail(&modelItems, customerID)
	return responses.NewResponseCart(c, http.StatusCreated, modelItems)
}

// Refresh godoc
// @Summary Read product count
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart/count [get]
func (h *HandlersShoppingCart) ReadItemCount(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelCount := models.CartItemCount{}
	cartRepo.ReadItemCount(&modelCount, customerID)
	return responses.NewResponseCartItemCount(c, http.StatusOK, modelCount)
}

// Refresh godoc
// @Summary Read cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [get]
func (h *HandlersShoppingCart) Read(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

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
// @Success 200 {object} responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart/{id} [put]
func (h *HandlersShoppingCart) UpdateQuantity(c echo.Context) error {
	cartID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	quantity, _ := strconv.ParseFloat(c.QueryParam("quantity"), 64)

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByCartID(&modelItem, cartID)
	if modelItem.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This cart item already exist in cart.")
	}

	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	prodRepo.ReadByID(&modelProduct, modelItem.VariationID)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.UpdateQuantity(cartID, &modelItem, modelProduct, quantity)
	modelItems := make([]models.CartItemsWithDetail, 0)
	cartRepo.ReadDetail(&modelItems, modelItem.CustomerID)
	return responses.NewResponseCart(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Delete cart item by ID
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart/{id} [delete]
func (h *HandlersShoppingCart) DeleteByID(c echo.Context) error {
	cartID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	cartRepo.ReadByCartID(&modelItem, cartID)

	if modelItem.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "No cart item exists at this ID.")
	}

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	cartService.Delete(cartID)
	return responses.MessageResponse(c, http.StatusOK, "The cart item successfully deleted.")
}

// Refresh godoc
// @Summary Delete all items in cart
// @Tags Shopping Cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/cart [delete]
func (h *HandlersShoppingCart) DeleteAll(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

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
