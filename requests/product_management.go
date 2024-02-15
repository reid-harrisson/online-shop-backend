package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProduct struct {
	StoreID          uint64   `json:"store_id" example:"1"`
	Name             string   `json:"name" example:"Honeycrisp"`
	ShortDescription string   `json:"short_description" example:"Apple"`
	LongDescirpiton  string   `json:"long_description" example:"Delicious Apple"`
	ImageUrls        []string `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
	UnitPriceRegular float64  `json:"unit_price_regular" example:"1.3"`
	StockQuantity    float64  `json:"stock_quantity" example:"47"`
	Active           int8     `json:"active" example:"1"`
}

type RequestProductChannel struct {
	Channels []string `json:"channels" example:"Factional (CPT),Finance,Coach Ally"`
}

type RequestProductContent struct {
	Contents []string `json:"contents" example:"My Life,Drumming Away,Winter Fashion"`
}

type RequestProductQuantity struct {
	Quantity float64 `json:"stock_quantity" example:"57"`
}

type RequestProductPrice struct {
	Price float64 `json:"price" example:"57"`
}

type RequestShippingData struct {
	Weight         float64 `json:"weight" example:"1.35"`
	Width          float64 `json:"width" example:"58"`
	Height         float64 `json:"height" example:"118"`
	Depth          float64 `json:"depth" example:"8"`
	Classification string  `json:"classification" example:"food"`
}

type RequestTag struct {
	Tags []string `json:"tags" example:"fruit,apple,agricultural products,grocery,food"`
}

type RequestAttributeItem struct {
	Name string `json:"name" example:"size"`
	Unit string `json:"unit" example:""`
}

type RequestAttribute struct {
	Attributes []RequestAttributeItem `json:"attributes"`
}

func (request RequestProduct) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.StoreID, validation.Required),
	)
}
