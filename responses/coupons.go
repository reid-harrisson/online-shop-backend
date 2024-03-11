package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type ResponseCoupon struct {
	ID                uint64    `json:"id"`
	StoreID           uint64    `json:"store_id"`
	CouponCode        string    `json:"coupon_code"`
	DiscountType      string    `json:"discount_type"`
	CouponAmount      float64   `json:"coupon_amount"`
	AllowFreeShipping int8      `json:"allow_free_shipping"`
	ExpiryDate        time.Time `json:"expiry_date"`
	MinimumSpend      float64   `json:"minimum_spend"`
	MaximumSpend      float64   `json:"maximum_spend"`
}

func NewResponseCoupon(c echo.Context, statusCode int, modelCoupon models.Coupons) error {
	return Response(c, statusCode, ResponseCoupon{
		ID:                uint64(modelCoupon.ID),
		StoreID:           modelCoupon.StoreID,
		CouponCode:        modelCoupon.CouponCode,
		DiscountType:      utils.CouponTypeToString(modelCoupon.DiscountType),
		CouponAmount:      modelCoupon.CouponAmount,
		AllowFreeShipping: modelCoupon.AllowFreeShipping,
		ExpiryDate:        modelCoupon.ExpiryDate,
		MinimumSpend:      modelCoupon.MinimumSpend,
		MaximumSpend:      modelCoupon.MaximumSpend,
	})
}

func NewResponseCoupons(c echo.Context, statusCode int, modelCoupons []models.Coupons) error {
	responseCoupon := []ResponseCoupon{}
	for _, modelCoupon := range modelCoupons {
		responseCoupon = append(responseCoupon, ResponseCoupon{
			ID:                uint64(modelCoupon.ID),
			StoreID:           modelCoupon.StoreID,
			CouponCode:        modelCoupon.CouponCode,
			DiscountType:      utils.CouponTypeToString(modelCoupon.DiscountType),
			CouponAmount:      modelCoupon.CouponAmount,
			AllowFreeShipping: modelCoupon.AllowFreeShipping,
			ExpiryDate:        modelCoupon.ExpiryDate,
			MinimumSpend:      modelCoupon.MinimumSpend,
			MaximumSpend:      modelCoupon.MaximumSpend,
		})
	}

	return Response(c, statusCode, responseCoupon)
}
