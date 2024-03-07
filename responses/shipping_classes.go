package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseShippingClass struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	StoreID     uint64 `json:"store_id"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}

func NewResponseShippingClass(c echo.Context, statusCode int, modelClass models.ShippingClasses) error {
	return Response(c, statusCode, ResponseShippingClass{
		ID:          uint64(modelClass.ID),
		Name:        modelClass.Name,
		StoreID:     modelClass.StoreID,
		Description: modelClass.Description,
		Priority:    modelClass.Priority,
	})
}
