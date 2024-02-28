package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseCartItem struct {
	ID             uint64              `json:"id"`
	ProductID      uint64              `json:"product_id"`
	ProductName    string              `json:"product_name"`
	ImageUrl       string              `json:"image_url"`
	Category       string              `json:"category"`
	Price          float64             `json:"price"`
	ExchangeRate   float64             `json:"exchange_rate"`
	ExchangePrice  float64             `json:"exchange_price"`
	DiscountAmount float64             `json:"discount_amount"`
	DiscountType   utils.DiscountTypes `json:"discount_type"`
	Quantity       float64             `json:"quantity"`
	TotalPrice     float64             `json:"total_price"`
}

type ResponseStoreCart struct {
	Items      []ResponseCartItem `json:"items"`
	StoreTotal float64            `json:"store_total"`
}

type ResponseCart struct {
	Stores     []ResponseStoreCart `json:"stores"`
	TotalPrice float64             `json:"total_price"`
}

type ResponseCartCount struct {
	Count uint64 `json:"count"`
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
			price := cartItem.Price
			switch cartItem.DiscountType {
			case utils.FixedAmountOff:
				price -= cartItem.DiscountAmount
			case utils.PercentageOff:
				price -= cartItem.DiscountAmount * price / 100
			}
			totalPrice := price * cartItem.Quantity
			responseCartItems = append(responseCartItems, ResponseCartItem{
				ID:          uint64(cartItem.ID),
				ProductID:   cartItem.VariationID,
				ProductName: cartItem.ProductName,
				ImageUrl:    imageUrl,
				Price:       price,
				Quantity:    cartItem.Quantity,
				Category:    cartItem.Category,
				TotalPrice:  totalPrice,
			})
			totalCost += totalPrice
		}
		responseStoreCarts = append(responseStoreCarts, ResponseStoreCart{
			Items:      responseCartItems,
			StoreTotal: totalCost,
		})
		totalPrice += totalCost
	}

	return Response(c, statusCode, ResponseCart{
		Stores:     responseStoreCarts,
		TotalPrice: totalPrice,
	})
}

func NewResponseCartItemCount(c echo.Context, statusCode int, modelCount models.CartItemCount) error {
	return Response(c, statusCode, ResponseCartCount{
		Count: modelCount.Count,
	})
}
