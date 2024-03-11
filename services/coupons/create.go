package cousvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"time"
)

func (service *Service) Create(modelCoupon *models.Coupons, req *requests.RequestCoupon, storeID uint64) error {
	modelCoupon.StoreID = storeID
	modelCoupon.CouponCode = req.CouponCode
	modelCoupon.DiscountType = req.DiscountType
	modelCoupon.CouponAmount = req.CouponAmount
	modelCoupon.AllowFreeShipping = req.AllowFreeShipping
	modelCoupon.ExpiryDate, _ = time.Parse("2006-01-02", req.ExpiryDate)
	modelCoupon.MinimumSpend = req.MinimumSpend
	modelCoupon.MaximumSpend = req.MaximumSpend
	return service.DB.Create(modelCoupon).Error
}
