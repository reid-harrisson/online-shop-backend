package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductCustomerRate struct {
	ID         uint64  `json:"id"`
	CustomerID uint64  `json:"customer_id"`
	ProductID  uint64  `json:"product_id"`
	Rate       float64 `json:"rate"`
}

type ResponseProductRate struct {
	ProductID uint64  `json:"product_id"`
	Customers uint64  `json:"customers"`
	Rate      float64 `json:"rate"`
}

func NewResponseProductCustomerRate(c echo.Context, statusCode int, modelRate models.ProductCustomerRates) error {
	responseRate := ResponseProductCustomerRate{
		ID:         uint64(modelRate.ID),
		CustomerID: modelRate.CustomerID,
		ProductID:  modelRate.ProductID,
		Rate:       modelRate.Rate,
	}
	return Response(c, statusCode, responseRate)
}

func NewResponseProductRate(c echo.Context, statusCode int, modelRate models.ProductRates) error {
	responseRate := ResponseProductRate{
		Customers: modelRate.Customers,
		ProductID: modelRate.ProductID,
		Rate:      modelRate.Rate,
	}
	return Response(c, statusCode, responseRate)
}
