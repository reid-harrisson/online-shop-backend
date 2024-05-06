package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProduct struct {
	StoreID          uint64   `json:"store_id" example:"1"`
	Title            string   `json:"name" example:"Honeycrisp"`
	ShortDescription string   `json:"short_description" example:"Apple"`
	LongDescription  string   `json:"long_description" example:"Delicious Apple"`
	ImageUrls        []string `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
}
type RequestProductWithDetail struct {
	RequestProduct
	RelatedChannels []uint64            `json:"channels" example:"241,249"`
	RelatedContents []uint64            `json:"contents" example:"549,552,558"`
	Categories      []uint64            `json:"categories" example:"4,3"`
	Tags            []string            `json:"tags" example:"fruit,food"`
	Attributes      map[string][]string `json:"attributes"`
	UpSell          []uint64            `json:"up_sell" example:"2,5,7"`
	CrossSell       []uint64            `json:"cross_sell" example:"3,6,8"`
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

func (request RequestProductWithDetail) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.StoreID, validation.Required),
	)
}
