package requests

type RequestProductTag struct {
	Tags []string `json:"tags" example:"Apple,Fruit,Food"`
}

type RequestTag struct {
	Name string `json:"tag" example:"Apple"`
}
