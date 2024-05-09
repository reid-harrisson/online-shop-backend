package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	eh "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/pkgs/utils"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cartsvc "OnlineStoreBackend/services/cart_items"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCart struct {
	server *s.Server
}

func NewHandlersCart(server *s.Server) *HandlersCart {
	return &HandlersCart{server: server}
}

// Refresh godoc
// @Summary Add cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param params body requests.RequestCartItem true "Variation Info"
// @Success 201 {object} responses.ResponseCartItem
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/cart [post]
func (h *HandlersCart) Create(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	req := new(requests.RequestCartItem)

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read product by id
	modelProduct := models.Products{}
	prodRepo := repositories.NewRepositoryProduct(h.server.DB)
	err = prodRepo.ReadByID(&modelProduct, req.ProductID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	if modelProduct.Status != utils.Approved {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.ProductNotApproved)
	}

	// Read variations by id
	modelVar := models.Variations{}
	modelVals := []models.AttributeValuesWithDetail{}
	valRepo := repositories.NewRepositoryAttributeValue(h.server.DB)
	err = valRepo.ReadByIDs(&modelVals, req.ValueIDs, req.ProductID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	mapVal := map[uint64]string{}
	for _, modelVal := range modelVals {
		if mapVal[modelVal.AttributeID] == "" {
			mapVal[modelVal.AttributeID] = modelVal.AttributeValue
		} else {
			return responses.ErrorResponse(c, http.StatusBadRequest, constants.DuplicateAttribute)
		}
	}

	// Read attribute value ids
	varRepo := repositories.NewRepositoryVariation(h.server.DB)
	err = varRepo.ReadByAttributeValueIDs(&modelVar, req.ValueIDs, req.ProductID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	variationID := uint64(modelVar.ID)
	log.Println("variationID = ", variationID)

	// Read cart item by info
	// Check duplicated cart
	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadByInfo(&modelItem, variationID, userID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode == 0 && message == "" {
		return responses.ErrorResponse(c, statusCode, constants.DuplicateCart)
	}

	// Create cart
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	err = cartService.Create(&modelItem, userID, &modelVar, float64(req.Quantity))
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCartItem(c, http.StatusCreated, modelItem)
}

// Refresh godoc
// @Summary Read count
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.ResponseCartCount
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// /@Router /store/api/v1/cart/count [get]
func (h *HandlersCart) ReadCount(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var count = int64(0)

	// Read item count
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadItemCount(&count, customerID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCartCount(c, http.StatusOK, count)
}

// Refresh godoc
// @Summary Read cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/cart [get]
func (h *HandlersCart) Read(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read cart detail
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	modelItems := make([]models.CartItemsWithDetail, 0)
	err = cartRepo.ReadDetail(&modelItems, customerID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCart(c, http.StatusOK, modelItems)
}

// Refresh godoc
// @Summary Edit cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Param quantity query string true "Quantity"
// @Success 200 {object} responses.ResponseCart
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/cart/{id} [put]
func (h *HandlersCart) UpdateQuantity(c echo.Context) error {
	cartID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}
	quantity, err := strconv.ParseFloat(c.QueryParam("quantity"), 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	var modelItem = models.CartItems{}

	// Update quqntity
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	err = cartService.UpdateQuantity(cartID, &modelItem, quantity)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCartItem(c, http.StatusOK, modelItem)
}

// Refresh godoc
// @Summary Delete a item
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Cart ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/cart/{id} [delete]
func (h *HandlersCart) Delete(c echo.Context) error {
	cartID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read cart by id
	modelItem := models.CartItems{}
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadByID(&modelItem, cartID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Delete cart
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	err = cartService.Delete(cartID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.DeleteCart)
}

// Refresh godoc
// @Summary Delete all items in cart
// @Tags Cart Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/cart [delete]
func (h *HandlersCart) DeleteAll(c echo.Context) error {
	customerID, err := strconv.ParseUint(c.Request().Header.Get("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read cart by customer id
	modelItems := make([]models.CartItems, 0)
	cartRepo := repositories.NewRepositoryCart(h.server.DB)
	err = cartRepo.ReadByCustomerID(&modelItems, customerID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	// Delete all carts
	cartService := cartsvc.NewServiceCartItem(h.server.DB)
	err = cartService.DeleteAll(customerID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.DeleteAllCart)
}
