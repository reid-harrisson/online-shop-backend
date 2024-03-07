package requests

type RequestShippingZone struct {
	Name        string   `json:"name" example:"Africa"`
	Places      []string `json:"regions" example:"South Africa,Johannesburg,London,Warsaw"`
	Description string   `json:"description" example:"example description"`
}
