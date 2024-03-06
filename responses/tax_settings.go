package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseTaxSetting struct {
	CustomerID uint64  `json:"customer_id"`
	CountryID  uint64  `json:"country_id"`
	TaxRate    float64 `json:"tax_rate"`
}

func NewResponseTaxSetting(c echo.Context, statusCode int, modelTax models.TaxSettings) error {
	responseTax := ResponseTaxSetting{
		CustomerID: modelTax.CustomerID,
		CountryID:  modelTax.CountryID,
		TaxRate:    modelTax.TaxRate,
	}

	return Response(c, statusCode, responseTax)
}
