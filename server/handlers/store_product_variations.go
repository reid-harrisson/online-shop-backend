package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	varsvc "OnlineStoreBackend/services/store_product_variations"
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
// @Summary Create prodcut variation
// @Tags Product Variation
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param product_id query int true "Product ID"
// @Param params body requests.RequestStoreProductVariation true "Variation Infomratiom"
// @Success 201 {object} responses.ResponseProductVariation
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/variation [post]
func (h *HandlersProductVariations) Create(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.QueryParam("product_id"), 10, 64)
	req := new(requests.RequestStoreProductVariation)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelVar := models.StoreProductVariations{}
	varService := varsvc.NewServiceProductVariation(h.server.DB)
	varService.Create(&modelVar, req, productID)

	return responses.NewResponseStoreProductVariation(c, http.StatusCreated, modelVar)
}
