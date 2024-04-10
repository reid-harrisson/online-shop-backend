package handlers

import (
	"OnlineStoreBackend/models"
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
// @Router /store/api/v1/coupon [post]
func (h *HandlersCoupons) Create(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)
	req := new(requests.RequestCoupon)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if req.CouponCode == "" {
		req.CouponCode = randomString(10)
	}

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	if err := couRepo.ReadByCode(&modelCoupon, req.CouponCode); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This coupon code already exist.")
	}

	couService := cousvc.NewServiceCoupon(h.server.DB)
	couService.Create(&modelCoupon, req, storeID)

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
// @Router /store/api/v1/coupon [get]
func (h *HandlersCoupons) Read(c echo.Context) error {
	storeID, _ := strconv.ParseUint(c.QueryParam("store_id"), 10, 64)

	modelCoupons := []models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	couRepo.ReadByStoreID(&modelCoupons, storeID)

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
// @Router /store/api/v1/coupon/{id} [put]
func (h *HandlersCoupons) Update(c echo.Context) error {
	couponID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	req := new(requests.RequestCoupon)
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	if err := couRepo.ReadByID(&modelCoupon, couponID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This coupon doesn't exist.")
	}
	if err := couRepo.ReadByCode(&modelCoupon, req.CouponCode); err == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This coupon code already exist.")
	}
	couService := cousvc.NewServiceCoupon(h.server.DB)
	if err := couService.Update(&modelCoupon, req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Fail to delete coupon.")
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
// @Router /store/api/v1/coupon/{id} [delete]
func (h *HandlersCoupons) Delete(c echo.Context) error {
	couponID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	modelCoupon := models.Coupons{}
	couRepo := repositories.NewRepositoryCoupon(h.server.DB)
	if err := couRepo.ReadByID(&modelCoupon, couponID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "This coupon doesn't exist.")
	}
	couService := cousvc.NewServiceCoupon(h.server.DB)
	if err := couService.Delete(couponID); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Fail to delete coupon.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Coupon successfully deleted.")
}
