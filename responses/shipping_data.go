package responses

type ResponseShipping struct {
	Weight         float64 `json:"weight"`
	Dimension      string  `json:"dimension"`
	Classification string  `json:"classification"`
}
