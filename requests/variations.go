package requests

type RequestProductVariation struct {
	Variants []string `json:"variants" example:"100,200,300"`
}
