package requests

type RequestCombo struct {
	DiscountAmount float64            `json:"discount_amount" example:"12"`
	DiscountType   string             `json:"discount_type" example:"Percentage Off"`
	ImageUrls      []string           `json:"image_urls" example:"https://www.pockittv.com/images/companies/63/products/bg_645a225496947_stirrup-cover-red-brass.webp,https://www.pockittv.com/images/companies/63/products/bg_645a339f5bef2_tall-black.webp"`
	Description    string             `json:"description" example:"Example description."`
	Title          string             `json:"title" example:"Example Title"`
	Items          []RequestComboItem `json:"items"`
}

type RequestComboItem struct {
	VariationID uint64  `json:"variation_id" example:"60"`
	Quantity    float64 `json:"quantity" example:"1"`
}
