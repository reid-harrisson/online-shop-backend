package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseCartItem struct {
	ID         uint64  `json:"id"`
	StoreID    uint64  `json:"store_id"`
	CustomerID uint64  `json:"customer_id"`
	ProductID  uint64  `json:"product_id"`
	Quantity   float64 `json:"quantity"`
}

type ResponseCartItemWithPrice struct {
	ResponseCartItem
	UnitPriceSale float64 `json:"unit_price_sale"`
	Price         float64 `json:"price"`
}

type ResponseCartWithTotal struct {
	Items    []ResponseCartItemWithPrice `json:"items"`
	SubTotal float64                     `json:"sub_total"`
}

type ResponseCartWithTax struct {
	Items     []ResponseCartItemWithPrice `json:"items"`
	SubTotal  float64                     `json:"sub_total"`
	TaxRate   float64                     `json:"tax_rate"`
	TaxAmount float64                     `json:"tax_amount"`
	Total     float64                     `json:"total"`
}

func NewResponseCartItem(c echo.Context, statusCode int, modelCart models.CartItems) error {
	responseCartItem := ResponseCartItem{
		ID:         uint64(modelCart.ID),
		StoreID:    modelCart.StoreID,
		ProductID:  modelCart.ProductID,
		CustomerID: modelCart.CustomerID,
		Quantity:   modelCart.Quantity,
	}
	return Response(c, statusCode, responseCartItem)
}

func NewResponseCarts(c echo.Context, statusCode int, modelCarts []models.CartItemWithPrice) error {
	responseCartItems := make([]ResponseCartItemWithPrice, 0)
	total := float64(0.0)
	for _, modelCart := range modelCarts {
		responseCartItems = append(responseCartItems, ResponseCartItemWithPrice{
			ResponseCartItem: ResponseCartItem{
				ID:         uint64(modelCart.ID),
				StoreID:    modelCart.StoreID,
				ProductID:  modelCart.ProductID,
				CustomerID: modelCart.CustomerID,
				Quantity:   modelCart.Quantity,
			},
			Price:         modelCart.Price,
			UnitPriceSale: modelCart.UnitPriceSale,
		})
		total += modelCart.Price
	}
	return Response(c, statusCode, ResponseCartWithTotal{
		SubTotal: total,
		Items:    responseCartItems,
	})
}

func NewResponseCartsPreview(c echo.Context, statusCode int, modelCarts []models.CartItemWithPrice, modelTaxSet models.TaxSettings) error {
	responseCartItems := make([]ResponseCartItemWithPrice, 0)
	total := float64(0.0)
	for _, modelCart := range modelCarts {
		responseCartItems = append(responseCartItems, ResponseCartItemWithPrice{
			ResponseCartItem: ResponseCartItem{
				ID:         uint64(modelCart.ID),
				StoreID:    modelCart.StoreID,
				ProductID:  modelCart.ProductID,
				CustomerID: modelCart.CustomerID,
				Quantity:   modelCart.Quantity,
			},
			Price:         modelCart.Price,
			UnitPriceSale: modelCart.UnitPriceSale,
		})
		total += modelCart.Price
	}
	return Response(c, statusCode, ResponseCartWithTax{
		SubTotal:  total,
		TaxRate:   modelTaxSet.TaxRate,
		TaxAmount: modelTaxSet.TaxRate * total / 100,
		Total:     total + modelTaxSet.TaxRate*total/100,
		Items:     responseCartItems,
	})
}
