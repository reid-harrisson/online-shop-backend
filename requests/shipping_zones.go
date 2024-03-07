package requests

type RequestShippingZone struct {
	Places      []string `json:"regions" example:"South Africa, Johannesburg, London, Warsaw"`
	Description string   `json:"description" example:"example description"`
}
