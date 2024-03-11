package responses

import (
	"OnlineStoreBackend/models"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ResponseShippingPlace struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type ResponseShippingZone struct {
	ID          uint64                  `json:"id"`
	Name        string                  `json:"name"`
	StoreID     uint64                  `json:"store_id"`
	Description string                  `json:"description"`
	Places      []ResponseShippingPlace `json:"places"`
}

func NewResponseShippingZone(c echo.Context, statusCode int, modelZone models.ShippingZonesWithPlace) error {
	responsePlaces := []ResponseShippingPlace{}
	placeIDs := strings.Split(modelZone.PlaceIDs, ",")
	placeNames := strings.Split(modelZone.PlaceNames, ",")
	for index := range placeIDs {
		placeID, _ := strconv.ParseUint(placeIDs[index], 10, 64)
		responsePlaces = append(responsePlaces, ResponseShippingPlace{
			ID:   placeID,
			Name: placeNames[index],
		})
	}
	return Response(c, statusCode, ResponseShippingZone{
		ID:          uint64(modelZone.ID),
		Name:        modelZone.Name,
		StoreID:     modelZone.StoreID,
		Description: modelZone.Description,
		Places:      responsePlaces,
	})
}
