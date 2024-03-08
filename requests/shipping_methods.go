package requests

import (
	"OnlineStoreBackend/pkgs/utils"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestShippingLocalPickup struct {
	Title       string  `json:"title" example:"Local pickup"`
	Description string  `json:"description" example:"Allow customers to pick up orders themselves. By default, when using local pickup store base taxes will apply regardless of customer address."`
	ZoneID      uint64  `json:"zone_id" example:"2"`
	TaxStatus   int8    `json:"tax_status" example:"1"`
	Cost        float64 `json:"cost" example:"2"`
}

type RequestShippingFree struct {
	Title              string             `json:"title" example:"Free shipping"`
	Description        string             `json:"description" example:"Free shipping is a special method which can be triggered with coupons and minimum spends."`
	ZoneID             uint64             `json:"zone_id" example:"2"`
	Requirement        utils.Requirements `json:"requirement" example:"2"`
	MinimumOrderAmount float64            `json:"minimum_order_amount" example:"2"`
}

type RequestFlatRate struct {
	ClassID     uint64  `json:"class_id" example:"1"`
	BaseCost    float64 `json:"base_cost" example:"2"`
	CostPerItem float64 `json:"cost_per_item" example:"2"`
	Percent     float64 `json:"percent" example:"0"`
	MinFee      float64 `json:"min_fee" example:"0"`
	MaxFee      float64 `json:"max_fee" example:"0"`
}
type RequestShippingFlatRate struct {
	Title       string            `json:"title" example:"Flat rate"`
	Description string            `json:"description" example:"Lets you charge a fixed rate for shipping."`
	ZoneID      uint64            `json:"zone_id" example:"2"`
	TaxStatus   int8              `json:"tax_status" example:"0"`
	Cost        float64           `json:"cost" example:"2"`
	Rates       []RequestFlatRate `json:"rates"`
}

func (request RequestShippingLocalPickup) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ZoneID, validation.Required),
		validation.Field(&request.TaxStatus, validation.Required),
		validation.Field(&request.Cost, validation.Required),
	)
}

func (request RequestShippingFree) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ZoneID, validation.Required),
		validation.Field(&request.Requirement, validation.Required),
		validation.Field(&request.MinimumOrderAmount, validation.Required),
	)
}

func (request RequestShippingFlatRate) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.ZoneID, validation.Required),
		validation.Field(&request.TaxStatus, validation.Required),
		validation.Field(&request.Cost, validation.Required),
	)
}
