package responses

import (
	"github.com/labstack/echo/v4"
)

type ResponseCheckoutVariation struct {
	ID            uint64   `json:"id"`
	VariationID   uint64   `json:"variation_id"`
	VariationName string   `json:"variation_name"`
	ImageUrls     []string `json:"image_urls"`
	Categories    []string `json:"categories"`
	SalePrice     float64  `json:"sale_price"`
	RegularPrice  float64  `json:"regular_price"`
	Quantity      float64  `json:"quantity"`
	StockLevel    float64  `json:"stock_level"`
	TotalPrice    float64  `json:"total_price"`
}

type ResponseCheckoutStore struct {
	Variations    []ResponseCheckoutVariation `json:"variations"`
	StoreID       uint64                      `json:"store_id"`
	SubTotal      float64                     `json:"sub_total"`
	ShippingPrice float64                     `json:"shipping_price"`
	TaxRate       float64                     `json:"tax_rate"`
	TaxAmount     float64                     `json:"tax_amount"`
	TotalPrice    float64                     `json:"total_price"`
}

type ResponseCheckout struct {
	Stores     []ResponseCheckoutStore `json:"stores"`
	TotalPrice float64                 `json:"total_price"`
}

func NewResponseCheckout(c echo.Context, statusCode int, responseStores []ResponseCheckoutStore) error {
	totalPrice := float64(0)
	for _, responseStore := range responseStores {
		totalPrice += responseStore.TotalPrice
	}

	return Response(c, statusCode, ResponseCheckout{
		Stores:     responseStores,
		TotalPrice: totalPrice,
	})
}
