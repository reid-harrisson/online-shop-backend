package requests

type RequestCheckout struct {
	CustomerID        uint64   `json:"customer_id" example:"1"`
	BillingAddressID  uint64   `json:"billing_address_id" example:"1"`
	ShippingAddressID uint64   `json:"shipping_address_id" example:"1"`
	CouponIDs         []uint64 `json:"coupon_code" example:"123"`
}
