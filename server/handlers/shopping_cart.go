package handlers

import (
	"OnlineStoreBackend/models"
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

func NewHandlersShoppingCart(server *s.Server) *HandlersShoppingCart {
	return &HandlersShoppingCart{server: server}
}

// Refresh godoc
// @Summary Add cart
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCartItem true "Cart Info"
// @Success 201 {object} responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [post]
func (h *HandlersShoppingCart) Create(c echo.Context) error {
	req := new(requests.RequestCartItem)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCart := models.CartItems{}
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	if err := cartService.Create(&modelCart, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseCartItem(c, http.StatusCreated, modelCart)
}

// Refresh godoc
// @Summary View cart
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Param store_id query int false "Store ID"
// @Success 200 {object} []responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [get]
func (h *HandlersShoppingCart) ReadAll(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelCarts := make([]models.CartItemWithPrice, 0)
	if err := cartRepo.ReadAll(&modelCarts, customerID, storeID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseCarts(c, http.StatusCreated, modelCarts)
}

// Refresh godoc
// @Summary View cart
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} []responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/preview [get]
func (h *HandlersShoppingCart) ReadPreview(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelCarts := make([]models.CartItemWithPrice, 0)
	modelTaxSet := models.TaxSettings{}
	if err := cartRepo.ReadPreview(&modelCarts, &modelTaxSet, customerID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseCartsPreview(c, http.StatusCreated, modelCarts, modelTaxSet)
}

// Refresh godoc
// @Summary Edit cart
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Param params body requests.RequestProductQuantity true "Product Quantity"
// @Success 200 {object} responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/{id} [put]
func (h *HandlersShoppingCart) UpdateQuantity(c echo.Context) error {
	req := new(requests.RequestMinimumStockLevel)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCart := models.CartItems{}
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	if err := cartService.UpdateQuantity(id, &modelCart, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.NewResponseCartItem(c, http.StatusCreated, modelCart)
}

// Refresh godoc
// @Summary Remove cart by Cart ID
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/{id} [delete]
func (h *HandlersShoppingCart) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	if err := cartService.Delete(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusCreated, "Successfully deleted")
}

// Refresh godoc
// @Summary Remove cart
// @Tags shopping cart
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param customer_id query int true "Customer ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [delete]
func (h *HandlersShoppingCart) DeleteAll(c echo.Context) error {
	customerID, _ := strconv.ParseUint(c.QueryParam("customer_id"), 10, 64)

	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	if err := cartService.DeleteAll(customerID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	return responses.MessageResponse(c, http.StatusCreated, "Successfully deleted")
}
