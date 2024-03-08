package responses

type ResponseFlatRate struct {
	ID          uint64  `json:"id"`
	ClassID     uint64  `json:"class_id"`
	BaseCost    float64 `json:"base_cost"`
	CostPerItem float64 `json:"cost_per_item"`
	Percent     float64 `json:"percent"`
	MinFee      float64 `json:"min_fee"`
	MaxFee      float64 `json:"max_fee"`
	ClassName   string  `json:"class_name"`
}
