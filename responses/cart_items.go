package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseProductCartItem struct {
	ProductID uint64  `json:"product_id"`
	CartID    uint64  `json:"cart_id"`
	Name      string  `json:"name"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
}

type ResponseStoreCartItem struct {
	Name     string                    `json:"name"`
	ID       uint64                    `json:"id"`
	Price    float64                   `json:"price"`
	Products []ResponseProductCartItem `json:"products"`
}

type ResponseCart struct {
	Total  float64                 `json:"total"`
	Stores []ResponseStoreCartItem `json:"stores"`
}

func NewResponseCart(c echo.Context, statusCode int, modelDetails []models.CartItemDetails) error {
	responseStores := make([]ResponseStoreCartItem, 0)
	index := 0
	total := 0.0
	price := 0.0
	for _, model := range modelDetails {
		if index == 0 || responseStores[index-1].Name != model.Store {
			if index != 0 {
				responseStores[index-1].Price = price
				total += price
			}
			responseStores = append(responseStores, ResponseStoreCartItem{
				Name:     model.Store,
				ID:       model.StoreID,
				Products: make([]ResponseProductCartItem, 0),
			})
			index += 1
			price = 0.0
		}
		responseStores[index-1].Products = append(responseStores[index-1].Products, ResponseProductCartItem{
			ProductID: model.ProductID,
			CartID:    uint64(model.ID),
			Name:      model.Product,
			Quantity:  model.Quantity,
			Price:     model.ProductPrice,
		})
		price += model.ProductPrice * model.Quantity
	}
	if index != 0 {
		responseStores[index-1].Price = price
		total += price
	}
	return Response(c, statusCode, ResponseCart{
		Total:  total,
		Stores: responseStores,
	})
}
