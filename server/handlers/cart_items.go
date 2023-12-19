package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cart "OnlineStoreBackend/services/cart_items"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCartItem struct {
	server *s.Server
}

func NewHandlersCartItem(server *s.Server) *HandlersCartItem {
	return &HandlersCartItem{server: server}
}

// Refresh godoc
// @Summary Add Item to Cart
// @Description Perform add item to cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCartItem true "Item Info"
// @Success 201 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [post]
func (h *HandlersCartItem) Create(c echo.Context) error {
	request := new(requests.RequestCartItem)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := request.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "All fields are required.")
	}

	model := models.CartItems{}
	service := cart.NewServiceCartItem(h.server.DB)
	if err := service.Create(&model, request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to add item to cart.")
	} else if model.Quantity == 0 {
		service.Delete(uint64(model.ID))
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to add item to cart.")
	}

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, model.UserID)
	return responses.NewResponseCart(c, http.StatusCreated, modelDetails)
}

// Refresh godoc
// @Summary  Read items from Cart
// @Description Perform read items from cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id query int true "User ID"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [get]
func (h *HandlersCartItem) Read(c echo.Context) error {
	userID, _ := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, userID)
	return responses.NewResponseCart(c, http.StatusOK, modelDetails)
}

// Refresh godoc
// @Summary Edit item in Cart
// @Description Perform edit item in cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Item ID"
// @Param params body requests.RequestCartItem true "Item Info"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/{id} [put]
func (h *HandlersCartItem) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	request := new(requests.RequestCartItem)
	if err := c.Bind(request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if err := request.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "All fields are required.")
	}

	model := models.CartItems{}
	service := cart.NewServiceCartItem(h.server.DB)
	if err := service.Update(&model, id, request); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to edit item in cart.")
	}

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, model.UserID)
	return responses.NewResponseCart(c, http.StatusOK, modelDetails)
}

// Refresh godoc
// @Summary  Remove item from Cart
// @Description Perform remove item from cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Item ID"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/{id} [delete]
func (h *HandlersCartItem) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	model := models.CartItems{}
	service := cart.NewServiceCartItem(h.server.DB)
	if err := service.Delete(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to add product to cart.")
	}

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, model.UserID)
	return responses.NewResponseCart(c, http.StatusOK, modelDetails)
}

// Refresh godoc
// @Summary  Clear All Items from Cart
// @Description Perform clear all items from cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id query int true "User ID"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart [delete]
func (h *HandlersCartItem) DeleteAll(c echo.Context) error {
	userID, _ := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)

	model := models.CartItems{}
	service := cart.NewServiceCartItem(h.server.DB)
	if err := service.DeleteAll(userID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to add product to cart.")
	}

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, model.UserID)
	return responses.NewResponseCart(c, http.StatusOK, modelDetails)
}

// Refresh godoc
// @Summary  Buy All Items in Cart
// @Description Perform buy all items in cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user_id query int true "User ID"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Router /api/v1/cart/buy [delete]
func (h *HandlersCartItem) DeleteBuy(c echo.Context) error {
	userID, _ := strconv.ParseUint(c.QueryParam("user_id"), 10, 64)

	model := models.CartItems{}
	service := cart.NewServiceCartItem(h.server.DB)
	if err := service.DeleteBuy(userID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to add product to cart.")
	}

	modelDetails := make([]models.CartItemDetails, 0)
	repository := repositories.NewRepositoryCartItem(h.server.DB)
	repository.Read(&modelDetails, model.UserID)
	return responses.NewResponseCart(c, http.StatusOK, modelDetails)
}
