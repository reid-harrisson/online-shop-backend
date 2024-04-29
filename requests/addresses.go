package requests

type RequestAddress struct {
	Name         string `json:"name" example:"Address 1"`
	AddressLine1 string `json:"address_line1" example:"2334 Devenish St"`
	AddressLine2 string `json:"address_line2" example:""`
	CityID       uint64 `json:"city_id" example:"2750264"`
	RegionID     uint64 `json:"region_id" example:"3558"`
	CountryID    uint64 `json:"country_id" example:"226"`
	PostalCode   string `json:"postal_code" example:"1401"`
	SubUrb       string `json:"suburb" example:"Honeydew"`
}
