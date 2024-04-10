package responses

import (
	"OnlineStoreBackend/models"
	ordsvc "OnlineStoreBackend/services/orders"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseCartVariation struct {
	ID            uint64   `json:"id"`
	VariationID   uint64   `json:"variation_id"`
	VariationName string   `json:"variation_name"`
	ImageUrls     []string `json:"image_urls"`
	Categories    []string `json:"categories"`
	SalePrice     float64  `json:"sale_price"`
	RegularPrice  float64  `json:"regular_price"`
	Quantity      float64  `json:"quantity"`
	TotalPrice    float64  `json:"total_price"`
}

type ResponseCartItem struct {
	ID          uint64  `json:"id"`
	CustomerID  uint64  `json:"customer_id"`
	VariationID uint64  `json:"variation_id"`
	Quantity    float64 `json:"quantity"`
}

type ResponseCartStore struct {
	Variations []ResponseCartVariation `json:"variations"`
	TotalPrice float64                 `json:"total_price"`
}

type ResponseCart struct {
	Stores     []ResponseCartStore `json:"stores"`
	TotalPrice float64             `json:"total_price"`
}

type ResponseCartCount struct {
	Count int64 `json:"count"`
}

func NewResponseCartItem(c echo.Context, statusCode int, modelItem models.CartItems) error {
	return Response(c, statusCode, ResponseCartItem{
		ID:          uint64(modelItem.ID),
		CustomerID:  modelItem.CustomerID,
		VariationID: modelItem.VariationID,
		Quantity:    modelItem.Quantity,
	})
}

func NewResponseCart(c echo.Context, statusCode int, modelCartItems []models.CartItemsWithDetail) error {
	allCartItems := make(map[uint64][]models.CartItemsWithDetail)
	responseStoreCarts := make([]ResponseCartStore, 0)

	for _, modelCartItem := range modelCartItems {
		allCartItems[modelCartItem.StoreID] = append(allCartItems[modelCartItem.StoreID], modelCartItem)
	}

	totalPrice := float64(0)
	for _, modelCartItems := range allCartItems {
		responseCartItems := make([]ResponseCartVariation, 0)
		storeTotal := float64(0)
		for _, cartItem := range modelCartItems {
			imageUrls := make([]string, 0)
			json.Unmarshal([]byte(cartItem.ImageUrls), &imageUrls)
			categories := make([]string, 0)
			json.Unmarshal([]byte("["+cartItem.Categories+"]"), &categories)
			regularPrice := cartItem.Price
			salePrice := ordsvc.GetSalePrice(cartItem)
			totalPrice := salePrice * cartItem.Quantity
			responseCartItems = append(responseCartItems, ResponseCartVariation{
				ID:            uint64(cartItem.ID),
				VariationID:   cartItem.VariationID,
				VariationName: cartItem.VariationName,
				ImageUrls:     imageUrls,
				RegularPrice:  regularPrice,
				SalePrice:     salePrice,
				Quantity:      cartItem.Quantity,
				Categories:    categories,
				TotalPrice:    totalPrice,
			})
			storeTotal += totalPrice
		}
		responseStoreCarts = append(responseStoreCarts, ResponseCartStore{
			Variations: responseCartItems,
			TotalPrice: storeTotal,
		})
		totalPrice += storeTotal
	}

	return Response(c, statusCode, ResponseCart{
		Stores:     responseStoreCarts,
		TotalPrice: totalPrice,
	})
}

func NewResponseCartCount(c echo.Context, statusCode int, count int64) error {
	return Response(c, statusCode, ResponseCartCount{
		Count: count,
	})
}
