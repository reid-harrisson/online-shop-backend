package requests

type RequestAttribute struct {
	Name string `json:"name" example:"length"`
	Unit string `json:"unit" example:"cm"`
}
