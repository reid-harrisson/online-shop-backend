package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseStore struct {
	ID                     uint64  `json:"id"`
	CompanyID              uint64  `json:"company_id"`
	UserID                 uint64  `json:"user_id"`
	ContactPhone           string  `json:"contact_phone"`
	ContactEmail           string  `json:"contact_email"`
	ShowStockQuantity      int8    `json:"show_stock_quantity"`
	ShowOutOfStockProducts int8    `json:"show_out_of_stock_products"`
	DeliveryPolicy         string  `json:"delivery_policy"`
	ReturnsPolicy          string  `json:"returns_policy"`
	Terms                  string  `json:"terms"`
	FlatRateShipping       float64 `json:"flat_rate_shipping"`
	BackOrder              int8    `json:"back_order"`
	Active                 int8    `json:"active"`
}

type ResponseStockTracking struct {
	StockTracking string `json:"stock_tracking"`
}

type ResponseBackOrder struct {
	BackOrder string `json:"back_order"`
}

func NewResponseStore(c echo.Context, statusCode int, modelStore models.Stores) error {
	responseStore := ResponseStore{
		ID:                     uint64(modelStore.ID),
		CompanyID:              modelStore.CompanyID,
		OwnerID:                modelStore.OwnerID,
		ContactPhone:           modelStore.ContactPhone,
		ContactEmail:           modelStore.ContactEmail,
		ShowStockQuantity:      modelStore.ShowStockQuantityStatus,
		ShowOutOfStockProducts: modelStore.ShowOutOfStockStatus,
		DeliveryPolicy:         modelStore.DeliveryPolicy,
		ReturnsPolicy:          modelStore.ReturnsPolicy,
		Terms:                  modelStore.Terms,
		FlatRateShipping:       modelStore.FlatRateShipping,
		BackOrder:              modelStore.ShowBackOrderStatus,
		Active:                 modelStore.Active,
	}
	return Response(c, statusCode, responseStore)
}

func NewResponseStockTracking(c echo.Context, statusCode int, stockTracking int8) error {
	if stockTracking == 0 {
		return Response(c, statusCode, ResponseStockTracking{
			StockTracking: "Disabled",
		})
	}
	return Response(c, statusCode, ResponseStockTracking{
		StockTracking: "Enabled",
	})
}

func NewResponseBackOrder(c echo.Context, statusCode int, backOrder int8) error {
	if backOrder == 0 {
		return Response(c, statusCode, ResponseStockTracking{
			StockTracking: "Disabled",
		})
	}
	return Response(c, statusCode, ResponseStockTracking{
		StockTracking: "Enabled",
	})
}
