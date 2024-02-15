package requests

type RequestProductChannel struct {
	Channels []string `json:"channels" example:"Factional (CPT),Finance,Coach Ally"`
}
