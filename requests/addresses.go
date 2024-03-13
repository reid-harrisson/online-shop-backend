package requests

type RequestAddress struct {
	AddressLine1 string `json:"address_line1" example:"Mozart"`
	AddressLine2 string `json:"address_line2" example:"Arizona"`
	SubUrb       string `json:"suburb" example:"Honeydew"`
	CountryID    uint64 `json:"country_id" example:"226"`
	RegionID     uint64 `json:"region_id" example:"3858"`
	CityID       uint64 `json:"city_id" example:"2744603"`
	PostalCode   string `json:"postal_code" example:"2000"`
}
