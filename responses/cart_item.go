package responses

import (
	"OnlineStoreBackend/models"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type ResponseCartItem struct {
	ID          uint64  `json:"id"`
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	ImageUrl    string  `json:"image_url"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    float64 `json:"quantity"`
	Category    string  `json:"category"`
	TotalPrice  float64 `json:"total_price"`
}

type ResponseStoreCart struct {
	Items     []ResponseCartItem `json:"items"`
	TotalCost float64            `json:"total_cost"`
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
				ID:          uint64(cartItem.ID),
				ProductID:   cartItem.ProductID,
				ProductName: cartItem.ProductName,
				ImageUrl:    imageUrl,
				UnitPrice:   cartItem.UnitPrice,
				Quantity:    cartItem.Quantity,
				Category:    cartItem.Category,
				TotalPrice:  cartItem.TotalPrice,
			})
			totalCost += cartItem.TotalPrice
		}
		responseStoreCarts = append(responseStoreCarts, ResponseStoreCart{
			Items:     responseCartItems,
			TotalCost: totalCost,
		})
		totalPrice += totalCost
	}

	return Response(c, statusCode, ResponseCart{
		Stores:     responseStoreCarts,
		TotalPrice: totalPrice,
	})
}

func NewResponseOrderPreview(c echo.Context, statusCode int, modelItems []models.CartItemsWithDetail, modelTax models.TaxSettings) error {
	allItems := make(map[uint64][]models.CartItemsWithDetail)
	responseStores := make([]ResponseStoreCartWithDetail, 0)

	for _, modelItem := range modelItems {
		allItems[modelItem.StoreID] = append(allItems[modelItem.StoreID], modelItem)
	}

	totalPrice := float64(0)
	for _, modelItems := range allItems {
		responseItems := make([]ResponseCartItem, 0)
		totalCost := float64(0)
		for _, modelItem := range modelItems {
			imageUrls := make([]string, 0)
			imageUrl := ""
			json.Unmarshal([]byte(modelItem.ImageUrl), &imageUrls)
			if len(imageUrls) > 0 {
				imageUrl = string(imageUrls[0])
			}
			responseItems = append(responseItems, ResponseCartItem{
				ID:          uint64(modelItem.ID),
				ProductID:   modelItem.ProductID,
				ProductName: modelItem.ProductName,
				ImageUrl:    imageUrl,
				UnitPrice:   modelItem.UnitPrice,
				Quantity:    modelItem.Quantity,
				Category:    modelItem.Category,
				TotalPrice:  modelItem.TotalPrice,
			})
			totalCost += modelItem.TotalPrice
		}
		taxes := totalCost * modelTax.TaxRate / 100
		responseStores = append(responseStores, ResponseStoreCartWithDetail{
			Items:        responseItems,
			TotalCost:    totalCost,
			Taxes:        taxes,
			ShippingCost: 0,
			TotalPrice:   totalCost + taxes,
		})
		totalPrice += totalCost + taxes
	}

	return Response(c, statusCode, ResponseOrderPreview{
		Stores:     responseStores,
		TotalPrice: totalPrice,
	})
}
