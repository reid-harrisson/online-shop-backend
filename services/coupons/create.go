package cousvc

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/requests"
	"fmt"
	"time"
)

func (service *Service) Create(modelCoupon *models.Coupons, req *requests.RequestCoupon, storeID uint64) error {
	modelCoupon.StoreID = storeID
	modelCoupon.CouponCode = req.CouponCode
	modelCoupon.DiscountType = req.DiscountType
	modelCoupon.CouponAmount = req.CouponAmount
	modelCoupon.AllowFreeShipping = req.AllowFreeShipping
	var err error
	modelCoupon.ExpiryDate, err = time.Parse("2006-01-02", req.ExpiryDate)
	fmt.Println(err)
	modelCoupon.MinimumSpend = req.MinimumSpend
	modelCoupon.MaximumSpend = req.MaximumSpend
	return service.DB.Create(modelCoupon).Error
}
