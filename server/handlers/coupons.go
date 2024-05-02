package handlers

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/constants"
	eh "OnlineStoreBackend/pkgs/error"
	"OnlineStoreBackend/repositories"
	"OnlineStoreBackend/requests"
	"OnlineStoreBackend/responses"
	s "OnlineStoreBackend/server"
	cousvc "OnlineStoreBackend/services/coupons"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type HandlersCoupons struct {
	server *s.Server
}

func NewHandlersCoupons(server *s.Server) *HandlersCoupons {
	return &HandlersCoupons{server: server}
}

var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Refresh godoc
// @Summary Create coupon
// @Tags Coupon Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Param params body requests.RequestCoupon true "Coupon"
// @Success 201 {object} responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/coupon [post]
func (h *HandlersCoupons) Create(c echo.Context) error {
	req := new(requests.RequestCoupon)

	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if req.CouponCode == "" {
		req.CouponCode = randomString(10)
	}

	// Read by code
	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	if err := couRepo.ReadByCode(&modelCoupon, req.CouponCode); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.CouponDuplicated)
	}

	// Create coupon
	couService := cousvc.NewServiceCoupon(h.server.DB)
	err = couService.Create(&modelCoupon, req, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCoupon(c, http.StatusCreated, modelCoupon)
}

// Refresh godoc
// @Summary Read coupon
// @Tags Coupon Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param store_id query int true "Store ID"
// @Success 200 {object} []responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/coupon [get]
func (h *HandlersCoupons) Read(c echo.Context) error {
	storeID, err := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read coupon by id
	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	err = couRepo.ReadByStoreID(&modelCoupons, storeID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCoupons(c, http.StatusOK, modelCoupons)
}

// Refresh godoc
// @Summary Update coupon
// @Tags Coupon Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Coupon ID"
// @Param params body requests.RequestCoupon true "Coupon"
// @Success 200 {object} responses.ResponseCoupon
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/coupon/{id} [put]
func (h *HandlersCoupons) Update(c echo.Context) error {
	req := new(requests.RequestCoupon)

	couponID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)

	// Read coupon by id
	if err := couRepo.ReadByID(&modelCoupon, couponID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.CouponNotFound)
	}

	// Read code
	if err := couRepo.ReadByCode(&modelCoupon, req.CouponCode); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.CouponDuplicated)
	}

	// Update coupon
	couService := cousvc.NewServiceCoupon(h.server.DB)
	err = couService.Update(&modelCoupon, req)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.NewResponseCoupon(c, http.StatusOK, modelCoupon)
}

// Refresh godoc
// @Summary Delete coupon
// @Tags Coupon Actions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Coupon ID"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /store/api/v1/coupon/{id} [delete]
func (h *HandlersCoupons) Delete(c echo.Context) error {
	couponID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.InvalidData)
	}

	// Read coupon by id
	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	if err := couRepo.ReadByID(&modelCoupon, couponID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, constants.CouponNotFound)
	}

	// Delete coupon by id
	couService := cousvc.NewServiceCoupon(h.server.DB)
	err = couService.Delete(couponID)
	if statusCode, message := eh.SqlErrorHandler(err); statusCode != 0 {
		return responses.ErrorResponse(c, statusCode, message)
	}

	return responses.MessageResponse(c, http.StatusOK, constants.CouponDeleted)
}
