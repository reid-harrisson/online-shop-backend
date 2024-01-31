package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestProduct struct {
	StoreID          uint64   `json:"store_id" example:"1"`
	Name             string   `json:"name" example:"Honeycrisp"`
	Brief            string   `json:"brief" example:"Apple"`
	Description      string   `json:"description" example:"Delicious Apple"`
	ImageUrls        []string `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
	UnitPriceRegular float64  `json:"unit_price_regular" example:"1.3"`
	UnitPriceSale    float64  `json:"unit_price_sale" example:"1.3"`
	StockQuantity    float64  `json:"stock_quantity" example:"47"`
	Active           int8     `json:"active" example:"1"`
}

type RequestProductChannel struct {
	Channels []string `json:"channels" example:"Factional (CPT),Finance,Coach Ally"`
}

type RequestProductContent struct {
	Contents []string `json:"contents" example:"My Life,Drumming Away,Winter Fashion"`
}

type RequestProductTag struct {
	Tags []string `json:"tags" example:"fruit,apple,agricultural products,grocery,food"`
}

type RequestProductAttribute struct {
	Attributes map[string]string `json:"attributes" example:"size:medium,colour:red,flavour:good,unit:kg"`
}

type RequestProductLinked struct {
	LinkedProductIDs []uint64 `json:"linked_product_ids" example:"1"`
}

type RequestProductQuantity struct {
	Quantity float64 `json:"stock_quantity" example:"57"`
}

type RequestProductPrice struct {
	Price float64 `json:"price" example:"57"`
}

type RequestProductReview struct {
	CustomerID uint64 `json:"customer_id" example:"1080"`
	Comment    string `json:"comment" example:"These are very good delicious apple but anyone can't eat them, because there are made of binary."`
}

type RequestShippingData struct {
	Weight         float64 `json:"weight" example:"1.35"`
	Dimension      string  `json:"dimension" example:"58mm*118mm*8mm"`
	Classification string  `json:"classification" example:"food"`
}

func (request RequestProduct) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.StoreID, validation.Required),
	)
}
