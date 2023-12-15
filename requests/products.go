package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProduct struct {
	CompanyID              uint64            `json:"company_id" example:"1"`
	UserID                 uint64            `json:"user_id" example:"1"`
	Name                   string            `json:"name" example:"Honeycrisp"`
	Brief                  string            `json:"brief" example:"Apple"`
	Description            string            `json:"description" example:"Delicious Apple"`
	SKU                    string            `json:"sku" example:"APPLE00"`
	Images                 []string          `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
	Tags                   string            `json:"catalogues" example:"fruit,apple,agricultural products,grocery,food"`
	Contents               string            `json:"contents" example:"My Life,Drumming Away,Winter Fashion"`
	Channels               string            `json:"channels" example:"Factional (CPT),Finance,Coach Ally"`
	UnitPriceRegular       float64           `json:"unit_price_regular" example:"1.3"`
	UnitPriceSale          float64           `json:"unit_price_sale" example:"2.1"`
	StockQuantity          float64           `json:"stock_quantity" example:"1987"`
	ShippingWeight         float64           `json:"shiping_weight" example:"20"`
	ShippingDimension      string            `json:"shiping_dimension" example:"500*600*400"`
	ShippingClassification string            `json:"shiping_classification" example:"Fruit"`
	Attributes             map[string]string `json:"" example:"Size:Medium,Colour:Red,Flavour:Good"`
	LinkedProducts         string            `json:"linked_products" example:""`
	Active                 int8              `json:"active" example:"1"`
}

func (request RequestProduct) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.CompanyID, validation.Required),
		validation.Field(&request.UserID, validation.Required),
	)
}
