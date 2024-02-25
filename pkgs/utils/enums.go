package utils

type ShippingMethods int8
type SimpleStatuses int8
type ProductStatus int8
type DiscountTypes int8
type SellTypes int8
type OrderStatuses int8

const (
	PercentageOff DiscountTypes = iota + 1
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
		return "PickUp"
	case FlatRate:
		return "FlatRate"
	case TableRate:
		return "TableRate"
	case RealTimeCarrierRate:
		return "RealTimeCarrierRate"
	case FreeShipping:
		return "FreeShipping"
	}

	return "PickUp"
}

func ShippingMethodsFromString(method string) ShippingMethods {
	switch method {
	case "PickUp":
		return PickUp
	case "FlatRate":
		return FlatRate
	case "TableRate":
		return TableRate
	case "RealTimeCarrierRate":
		return RealTimeCarrierRate
	case "FreeShipping":
		return FreeShipping
	}

	return PickUp
}
