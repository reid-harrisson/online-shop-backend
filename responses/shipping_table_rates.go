package responses

type ResponseTableRate struct {
	ID          uint64  `json:"id"`
	Condition   string  `json:"condition"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	RowCost     float64 `json:"row_cost"`
	ItemCost    float64 `json:"item_cost"`
	CostPerKg   float64 `json:"cost_per_kg"`
	PercentCost float64 `json:"percent_cost"`
}
