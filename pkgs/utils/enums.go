package utils

type DiscountTypes int8

const (
	PercentageOff DiscountTypes = iota + 1
	FixedAmountOff
	FreeShipping
)
