package responses

type ResponseTableRate struct {
	ID          uint64  `json:"id"`
	ClassID     uint64  `json:"class_id"`
	Condition   string  `json:"condition"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Break       int8    `json:"break"`
	Abort       int8    `json:"abort"`
	RowCost     float64 `json:"row_cost"`
	ItemCost    float64 `json:"item_cost"`
	CostPerKg   float64 `json:"cost_per_kg"`
	PercentCost float64 `json:"percent_cost"`
	ClassName   string  `json:"class_name"`
}
