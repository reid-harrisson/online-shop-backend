package responses

import (
	"OnlineStoreBackend/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseCartItem struct {
	ID                uint64  `json:"id"`
	ProductID         uint64  `json:"product_id"`
	ProductName       string  `json:"product_name"`
	ImageUrl          string  `json:"image_url"`
	UnitPriceSale     float64 `json:"unit_price_sale"`
	ExchangeRate      float64 `json:"exchange_rate"`
	UnitPriceExchange float64 `json:"unit_price_exchange"`
	Quantity          float64 `json:"quantity"`
	Category          string  `json:"category"`
	TotalPrice        float64 `json:"total_price"`
}

type ResponseStoreCart struct {
	Items    []ResponseCartItem `json:"items"`
	SubTotal float64            `json:"sub_total"`
}

type ResponseStoreCartWithDetail struct {
	Items        []ResponseCartItem `json:"items"`
	TotalCost    float64            `json:"total_cost"`
	Taxes        float64            `json:"taxes"`
	ShippingCost float64            `json:"shipping_cost"`
	TotalPrice   float64            `json:"total_price"`
}

type ResponseCart struct {
	Stores     []ResponseStoreCart `json:"stores"`
	TotalPrice float64             `json:"total_price"`
}

type ResponseOrderPreview struct {
	Stores     []ResponseStoreCartWithDetail `json:"stores"`
	TotalPrice float64                       `json:"total_price"`
}

func NewResponseCart(c echo.Context, statusCode int, modelCartItems []models.CartItemsWithDetail) error {
	allCartItems := make(map[uint64][]models.CartItemsWithDetail)
	responseStoreCarts := make([]ResponseStoreCart, 0)

	for _, modelCartItem := range modelCartItems {
		allCartItems[modelCartItem.StoreID] = append(allCartItems[modelCartItem.StoreID], modelCartItem)
	}

	totalPrice := float64(0)
	for _, modelCartItems := range allCartItems {
		responseCartItems := make([]ResponseCartItem, 0)
		totalCost := float64(0)
		for _, cartItem := range modelCartItems {
			imageUrls := make([]string, 0)
			imageUrl := ""
			json.Unmarshal([]byte(cartItem.ImageUrl), &imageUrls)
			if len(imageUrls) > 0 {
				imageUrl = string(imageUrls[0])
			}
			responseCartItems = append(responseCartItems, ResponseCartItem{
				ID:            uint64(cartItem.ID),
				ProductID:     cartItem.ProductID,
				ProductName:   cartItem.ProductName,
				ImageUrl:      imageUrl,
				UnitPriceSale: cartItem.UnitPrice,
				Quantity:      cartItem.Quantity,
				Category:      cartItem.Category,
				TotalPrice:    cartItem.TotalPrice,
			})
			totalCost += cartItem.TotalPrice
		}
		responseStoreCarts = append(responseStoreCarts, ResponseStoreCart{
			Items:    responseCartItems,
			SubTotal: totalCost,
		})
		totalPrice += totalCost
	}

	return Response(c, statusCode, ResponseCart{
		Stores:     responseStoreCarts,
		TotalPrice: totalPrice,
	})
}
