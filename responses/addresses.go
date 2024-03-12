package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseAddress struct {
	ID           uint64 `json:"id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	SubUrb       string `json:"suburb"`
	CountryID    uint64 `json:"country_id"`
	RegionID     uint64 `json:"region_id"`
	CityID       uint64 `json:"city_id"`
	PostalCode   string `json:"postal_code"`
}

func NewResponseAddress(c echo.Context, statusCode int, modelAddr models.Addresses) error {
	return Response(c, statusCode, ResponseAddress{
		ID:           uint64(modelAddr.ID),
		AddressLine1: modelAddr.AddressLine1,
		AddressLine2: modelAddr.AddressLine2,
		SubUrb:       modelAddr.SubUrb,
		CountryID:    modelAddr.CountryID,
		RegionID:     modelAddr.RegionID,
		CityID:       modelAddr.CityID,
		PostalCode:   modelAddr.PostalCode,
	})
}

func NewResponseAddresses(c echo.Context, statusCode int, modelAddrs []models.Addresses) error {
	responseAddrs := make([]ResponseAddress, 0)
	for _, modelAddr := range modelAddrs {
		responseAddrs = append(responseAddrs, ResponseAddress{
			ID:           uint64(modelAddr.ID),
			AddressLine1: modelAddr.AddressLine1,
			AddressLine2: modelAddr.AddressLine2,
			SubUrb:       modelAddr.SubUrb,
			CountryID:    modelAddr.CountryID,
			RegionID:     modelAddr.RegionID,
			CityID:       modelAddr.CityID,
			PostalCode:   modelAddr.PostalCode,
		})
	}
	return Response(c, statusCode, responseAddrs)
}
