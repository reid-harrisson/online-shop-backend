package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestShippingMethod struct {
	ZoneID              uint64             `json:"zone_id"`
	Method              string             `json:"method"`
	Requirement         utils.Requirements `json:"requirement"`
	MinimumOrderAmount  float64            `json:"minimum_order_amount"`
	TaxStatus           int8               `json:"tax_status"`
	Cost                float64            `json:"cost"`
	TaxIncluded         int8               `json:"tax_included"`
	HandlingFee         float64            `json:"handling_fee"`
	MaximumShippingCost float64            `json:"maximum_shipping_cost"`
	CalculationType     int8               `json:"calculation_type"`
	HandlingFeePerClass float64            `json:"handling_fee_per_class"`
	MinimumCostPerClass float64            `json:"minimum_cost_per_class"`
	MaximumCostPerClass float64            `json:"maximum_cost_per_class"`
	DiscountInMinMax    int8               `json:"discount_in_min_max"`
	TaxInMinMax         int8               `json:"tax_in_min_max"`
}

func (request RequestShippingMethod) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Method, validation.Required),
	)
}

type RequestShippingLocalPickup struct {
	ZoneID    uint64  `json:"zone_id"`
	TaxStatus int8    `json:"tax_status"`
	Cost      float64 `json:"cost"`
}

func (request RequestShippingLocalPickup) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ZoneID, validation.Required),
		validation.Field(&request.TaxStatus, validation.Required),
		validation.Field(&request.Cost, validation.Required),
	)
}

type RequestShippingFree struct {
	ZoneID             uint64             `json:"zone_id"`
	Requirement        utils.Requirements `json:"requirement"`
	MinimumOrderAmount float64            `json:"minimum_order_amount"`
}

func (request RequestShippingFree) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ZoneID, validation.Required),
		validation.Field(&request.Requirement, validation.Required),
		validation.Field(&request.MinimumOrderAmount, validation.Required),
	)
}
