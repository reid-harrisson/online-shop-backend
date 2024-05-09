package requests

type RequestCartItem struct {
	ProductID uint64   `json:"product_id" example:"32"`
	Quantity  float64  `json:"quantity" example:"2"`
	ValueIDs  []uint64 `json:"attribute_value_ids" example:"5"`
}
