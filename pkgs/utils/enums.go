package utils

type ShippingMethods int8
type ShowStockLevelStatus int8
type ShowOutOfStockStatus int8
type BackOrderStatus int8

const (
	PickUp ShippingMethods = iota
	FlatRate
	TableRate
	RealTimeCarrierRate
	FreeShipping
)

const (
	HideStockLevel ShowStockLevelStatus = iota
	ShowStockLevel
)

const (
	HideOutOfStock ShowOutOfStockStatus = iota
	ShowOutOfStock
)

const (
	HideBackOrder BackOrderStatus = iota
	ShowBackOrder
)
