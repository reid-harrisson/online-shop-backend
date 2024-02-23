package utils

type ShippingMethods int8
type StockLevelStatus int8
type OutOfStockStatus int8
type BackOrderStatus int8
type ProductStatus int8
type DiscountTypes int8
type SellTypes int8

const (
	UpSell SellTypes = iota
	CrossSell
)

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
	HideStockLevel StockLevelStatus = iota
	ShowStockLevel
)

const (
	HideOutOfStock OutOfStockStatus = iota
	ShowOutOfStock
)

const (
	HideBackOrder BackOrderStatus = iota
	ShowBackOrder
)

const (
	Draft ProductStatus = iota
	Pending
	Approved
	Rejected
)

func ProductStatusToString(productStatus ProductStatus) string {
	switch productStatus {
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

func StockLevelStatusToString(stockLevelStatus StockLevelStatus) string {
	switch stockLevelStatus {
	case HideStockLevel:
		return "Off"
	case ShowStockLevel:
		return "On"
	}

	return "Off"
}

func OutOfStockStatusToString(outOfStockStatus OutOfStockStatus) string {
	switch outOfStockStatus {
	case HideOutOfStock:
		return "Off"
	case ShowOutOfStock:
		return "On"
	}

	return "Off"
}

func BackOrderStatusToString(backOrderStatus BackOrderStatus) string {
	switch backOrderStatus {
	case HideBackOrder:
		return "Off"
	case ShowBackOrder:
		return "On"
	}

	return "Off"
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
