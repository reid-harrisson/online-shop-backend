package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseCartItem struct {
	ID           uint64   `json:"id"`
	ProductID    uint64   `json:"product_id"`
	ProductName  string   `json:"product_name"`
	ImageUrls    []string `json:"image_urls"`
	Categories   []string `json:"categories"`
	SalePrice    float64  `json:"sale_price"`
	RegularPrice float64  `json:"regular_price"`
	Quantity     float64  `json:"quantity"`
	TotalPrice   float64  `json:"total_price"`
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
			json.Unmarshal([]byte(cartItem.ImageUrls), &imageUrls)
			categories := make([]string, 0)
			json.Unmarshal([]byte("["+cartItem.Categories+"]"), &categories)
			regularPrice := cartItem.Price
			salePrice := cartItem.Price
			switch cartItem.DiscountType {
			case utils.FixedAmountOff:
				salePrice -= cartItem.DiscountAmount
			case utils.PercentageOff:
				salePrice -= cartItem.DiscountAmount * salePrice / 100
			}
			totalPrice := salePrice * cartItem.Quantity
			responseCartItems = append(responseCartItems, ResponseCartItem{
				ID:           uint64(cartItem.ID),
				ProductID:    cartItem.VariationID,
				ProductName:  cartItem.VariationName,
				ImageUrls:    imageUrls,
				RegularPrice: regularPrice,
				SalePrice:    salePrice,
				Quantity:     cartItem.Quantity,
				Categories:   categories,
				TotalPrice:   totalPrice,
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
