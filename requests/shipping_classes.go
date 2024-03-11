package requests

type RequestShippingClass struct {
	Name        string `json:"name" example:"Courier Refrigerated"`
	Description string `json:"description" example:"Shipping Classification For Food"`
	Priority    int8   `json:"priority" example:"2"`
}
