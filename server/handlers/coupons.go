package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cousvc "OnlineStoreBackend/services/coupons"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlersCoupons struct {
	server *s.Server
}

func NewHandlersCoupons(server *s.Server) *HandlersCoupons {
	return &HandlersCoupons{server: server}
}

// Refresh godoc
// @Summary Create coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCoupo true "Coupon"
// @Success 201 {object} responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/coupon [post]
func (h *HandlersCoupons) Create(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestCoupon)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCoupon := models.Coupons{}
	couService := cousvc.NewServiceCoupon(h.server.DB)
	couService.Create(&modelCoupon, req, storeID)

	return responses.NewResponseCoupon(c, http.StatusCreated, modelCoupon)
}

// Refresh godoc
// @Summary Read coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id path int true "Store ID"
// @Success 200 {object} []responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Router /store/api/v1/store/:id/coupon [post]
func (h *HandlersCoupons) Read(c echo.Context) error {

	modelCoupon := models.Coupons{}

	return responses.NewResponseCoupon(c, http.StatusCreated, modelCoupon)
}
