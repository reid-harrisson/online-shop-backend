package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProduct struct {
	StoreID          uint64   `json:"store_id" example:"1"`
	Title            string   `json:"name" example:"Honeycrisp"`
	ShortDescription string   `json:"short_description" example:"Apple"`
	LongDescirpiton  string   `json:"long_description" example:"Delicious Apple"`
	ImageUrls        []string `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
}

type RequestMinimumStockLevel struct {
	Level float64 `json:"level" example:"57"`
}

type RequestProductPrice struct {
	Price float64 `json:"price" example:"57"`
}

func (request RequestProduct) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.StoreID, validation.Required),
	)
}
