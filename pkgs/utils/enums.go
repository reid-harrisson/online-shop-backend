package utils

type ShippingMethods int8
type SimpleStatuses int8
type ProductStatus int8
type DiscountTypes int8
type SellTypes int8

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
