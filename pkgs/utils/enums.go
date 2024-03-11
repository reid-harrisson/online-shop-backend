package utils

type ShippingMethods int8
type SimpleStatuses int8
type ProductStatus int8
type DiscountTypes int8
type SellTypes int8
type OrderStatuses int8
type PageTypes int8
type ProductTypes int8
type Requirements int8
type Conditions int8
type CouponTypes int8

const (
	PercentageDiscount CouponTypes = iota
	FixedCartDiscount
	FixedProductDiscount
)

const (
	None Conditions = iota
	Price
	Weight
	ItemCount
	Width
	Length
	Height
	ItemCountSameClass
)

const (
	NoRequirement Requirements = iota
	ValidFreeShippingCoupon
	MinimumOrderAmount
	MinimumOrderAmountOrCoupon
	MinimumOrderAmountAndCoupon
)

const (
	Simple ProductTypes = iota
	Variable
)

const (
	StorePage PageTypes = iota
	ProductPage
	ProductDetailPage
	LoginPage
	RegisterPage
	CartPage
	PaymentPage
	PaymentConfirmPage
)

const (
	PercentageOff DiscountTypes = iota
	FixedAmountOff
)

const (
	PickUp ShippingMethods = iota
	FlatRate
	TableRate
	RealTimeCarrierRate
	FreeShipping
)

const (
	Disabled SimpleStatuses = iota
	Enabled
)

const (
	Draft ProductStatus = iota
	Pending
	Approved
	Rejected
)

const (
	UpSell SellTypes = iota
	CrossSell
)

const (
	StatusOrderPending OrderStatuses = iota + 1
	StatusOrderBackOrdered
	StatusOrderPaymentProcessing
	StatusOrderPaid
	StatusOrderProcessing
	StatusOrderShippingProcessing
	StatusOrderShipping
	StatusOrderShipped
	StatusOrderCompleted
)

type ReviewStatuses int8

const (
	StatusReviewPending ReviewStatuses = iota + 0
	StatusReviewApproved
	StatusReviewBlocked
)

func CouponTypeToString(couponType CouponTypes) string {
	switch couponType {
	case PercentageDiscount:
		return "Percentage discount"
	case FixedCartDiscount:
		return "Fixed cart discount"
	case FixedProductDiscount:
		return "Fixed product discount"
	}
	return "Percentage discount"
}

func CouponTypeFromString(couponType string) CouponTypes {
	switch couponType {
	case "Percentage discount":
		return PercentageDiscount
	case "Fixed cart discount":
		return FixedCartDiscount
	case "Fixed product discount":
		return FixedProductDiscount
	}
	return PercentageDiscount
}

func ConditionToString(condition Conditions) string {
	switch condition {
	case Price:
		return "Price"
	case Weight:
		return "Weight"
	case ItemCount:
		return "Item Count"
	case Width:
		return "Width"
	case Length:
		return "Length"
	case Height:
		return "Height"
	case ItemCountSameClass:
		return "Item Count (same class)"
	}
	return "None"
}

func ConditionFromString(condition string) Conditions {
	switch condition {
	case "Price":
		return Price
	case "Weight":
		return Weight
	case "Item Count":
		return ItemCount
	case "Width":
		return Width
	case "Length":
		return Length
	case "Height":
		return Height
	}
	return None
}

func RequirementToString(requirement Requirements) string {
	switch requirement {
	case ValidFreeShippingCoupon:
		return "A valid free shipping coupon"
	case MinimumOrderAmount:
		return "A minimum order amount"
	case MinimumOrderAmountOrCoupon:
		return "A minimum order amount OR coupon"
	case MinimumOrderAmountAndCoupon:
		return "A minimum order amount AND coupon"
	}
	return "No requirement"
}

func PageTypeFromString(pageType string) PageTypes {
	switch pageType {
	case "Store":
		return StorePage
	case "Product":
		return ProductPage
	case "Product Detail":
		return ProductDetailPage
	case "Login":
		return LoginPage
	case "Register":
		return RegisterPage
	case "Cart":
		return CartPage
	case "Payment":
		return PaymentPage
	case "Payment Confirm":
		return PaymentConfirmPage
	}
	return StorePage
}

func PageTypeToString(pageType PageTypes) string {
	switch pageType {
	case StorePage:
		return "Store Page"
	case ProductPage:
		return "Product Page"
	case ProductDetailPage:
		return "Product Detail Page"
	case LoginPage:
		return "Login Page"
	case RegisterPage:
		return "Register Page"
	case CartPage:
		return "Cart Page"
	case PaymentPage:
		return "Payment Page"
	case PaymentConfirmPage:
		return "Payment Confirm Page"
	}
	return "Store Page"
}

func ReviewStatusFromString(status string) ReviewStatuses {
	switch status {
	case "Pending":
		return StatusReviewPending
	case "Approved":
		return StatusReviewApproved
	case "Blocked":
		return StatusReviewBlocked
	}
	return StatusReviewPending
}

func OrderStatusToString(status OrderStatuses) string {
	switch status {
	case StatusOrderPending:
		return "Pending"
	case StatusOrderBackOrdered:
		return "Back Ordered"
	case StatusOrderPaymentProcessing:
		return "Payment Processing"
	case StatusOrderPaid:
		return "Paid"
	case StatusOrderProcessing:
		return "Processing"
	case StatusOrderShippingProcessing:
		return "Shipping Processing"
	case StatusOrderShipping:
		return "Shipping"
	case StatusOrderShipped:
		return "Shipped"
	case StatusOrderCompleted:
		return "Completed"
	}
	return "Pending"
}

func OrderStatusFromString(status string) OrderStatuses {
	switch status {
	case "Pending":
		return StatusOrderPending
	case "Payment Processing":
		return StatusOrderPaymentProcessing
	case "Paid":
		return StatusOrderPaid
	case "Processing":
		return StatusOrderProcessing
	case "Shipping Processing":
		return StatusOrderShippingProcessing
	case "Shipping":
		return StatusOrderShipping
	case "Shipped":
		return StatusOrderShipped
	case "Completed":
		return StatusOrderCompleted
	}
	return StatusOrderPending
}

func DiscountTypeToString(discountType DiscountTypes) string {
	switch discountType {
	case PercentageOff:
		return "Percentage Off"
	case FixedAmountOff:
		return "Fixed Amount Off"
	}
	return "Fixed Amount Off"
}

func ProductStatusToString(status ProductStatus) string {
	switch status {
	case Draft:
		return "Draft"
	case Pending:
		return "Pending"
	case Approved:
		return "Approved"
	case Rejected:
		return "Rejected"
	}

	return "Draft"
}

func ProductStatusFromString(productStatus string) ProductStatus {
	switch productStatus {
	case "Draft":
		return Draft
	case "Pending":
		return Pending
	case "Approved":
		return Approved
	case "Rejected":
		return Rejected
	}

	return Draft
}

func SimpleStatusToString(status SimpleStatuses) string {
	switch status {
	case Disabled:
		return "Disabled"
	case Enabled:
		return "Enabled"
	}

	return "Disabled"
}

func ShippingMethodsToString(method ShippingMethods) string {
	switch method {
	case PickUp:
		return "Local pickup"
	case FlatRate:
		return "Flat rate"
	case TableRate:
		return "Table rate"
	case RealTimeCarrierRate:
		return "Real time carrier rate"
	case FreeShipping:
		return "Free shipping"
	}
	return "Pick up"
}

func ShippingMethodsFromString(method string) ShippingMethods {
	switch method {
	case "Local pickup":
		return PickUp
	case "Flat rate":
		return FlatRate
	case "Table rate":
		return TableRate
	case "Real time carrier rate":
		return RealTimeCarrierRate
	case "Free shipping":
		return FreeShipping
	}
	return PickUp
}

func SellTypesFromString(sellType string) SellTypes {
	switch sellType {
	case "UpSell":
		return UpSell
	case "CrossSell":
		return CrossSell
	}

	return UpSell
}
