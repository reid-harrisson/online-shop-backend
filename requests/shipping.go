package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestShippingMethod struct {
	ZoneID              uint64  `json:"zone_id"`
	Method              string  `json:"method"`
	Requirement         int8    `json:"requirement"`
	MinimumOrderAmount  float64 `json:"minimum_order_amount"`
	TaxStatus           int8    `json:"tax_status"`
	Cost                float64 `json:"cost"`
	TaxIncluded         int8    `json:"tax_included"`
	HandlingFee         float64 `json:"handling_fee"`
	MaximumShippingCost float64 `json:"maximum_shipping_cost"`
	CalculationType     int8    `json:"calculation_type"`
	HandlingFeePerClass float64 `json:"handling_fee_per_class"`
	MinimumCostPerClass float64 `json:"minimum_cost_per_class"`
	MaximumCostPerClass float64 `json:"maximum_cost_per_class"`
	DiscountInMinMax    int8    `json:"discount_in_min_max"`
	TaxInMinMax         int8    `json:"tax_in_min_max"`
}

type RequestShippingData struct {
	Weight         float64 `json:"weight" example:"1.35"`
	Width          float64 `json:"width" example:"58"`
	Height         float64 `json:"height" example:"118"`
	Length         float64 `json:"length" example:"8"`
	Classification string  `json:"classification" example:"food"`
}

func (request RequestShippingMethod) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Method, validation.Required),
	)
}

func (request RequestShippingData) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Weight, validation.Required),
		validation.Field(&request.Width, validation.Required),
		validation.Field(&request.Height, validation.Required),
		validation.Field(&request.Length, validation.Required),
		validation.Field(&request.Classification, validation.Required),
	)
}
