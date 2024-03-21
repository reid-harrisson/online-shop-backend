package utils

type TrackEvents int8

const (
	OrderPlaced TrackEvents = iota + 1
	OrderCancelled
	ProductWarhousing
)

type PaymentType int8

const (
	WalletDeposit PaymentType = iota + 1
	WalletWithdraw
	StorePurchase
)

type PageTypes int8

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

type ProductTypes int8

const (
	Simple ProductTypes = iota
	Variable
)

type Requirements int8

const (
	NoRequirement Requirements = iota
	ValidFreeShippingCoupon
	MinimumOrderAmount
	MinimumOrderAmountOrCoupon
	MinimumOrderAmountAndCoupon
)

type Conditions int8

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

type CouponTypes int8

const (
	PercentageDiscount CouponTypes = iota
	FixedCartDiscount
	FixedProductDiscount
)

type DiscountTypes int8

const (
	PercentageOff DiscountTypes = iota
	FixedAmountOff
)

type ShippingMethods int8

const (
	PickUp ShippingMethods = iota
	FlatRate
	TableRate
	RealTimeCarrierRate
	FreeShipping
)

type SimpleStatuses int8

const (
	Disabled SimpleStatuses = iota
	Enabled
)

type ProductStatus int8

const (
	Draft ProductStatus = iota
	Pending
	Approved
	Rejected
)

type SellTypes int8

const (
	UpSell SellTypes = iota
	CrossSell
)

type OrderStatuses int8

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
	if discountType == PercentageOff {
		return "Percentage Off"
	}
	return "Fixed Amount Off"
}

func DiscountTypeFromString(discountType string) DiscountTypes {
	if discountType == "Percentage Off" {
		return PercentageOff
	}
	return FixedAmountOff

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

func ProductTypesFromString(productType string) ProductTypes {
	switch productType {
	case "simple":
		return Simple
	case "variable":
		return Variable
	}
	return Variable
}
