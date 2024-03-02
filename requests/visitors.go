package requests

type RequestVisitor struct {
	StoreID     uint64  `json:"store_id" example:"1"`
	ProductID   uint64  `json:"product_id" example:"31"`
	IpAddress   string  `json:"ip_address" example:"11.111.11.111"`
	Page        string  `json:"page" example:"Store"`
	Bounce      uint64  `json:"bounce" example:"1"`
	LoadingTime float64 `json:"loading_time" example:"0.2"`
}
