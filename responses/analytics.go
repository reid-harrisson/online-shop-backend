package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseSalesRevenue struct {
	StoreID uint64  `json:"store_id"`
	Revenue float64 `json:"revenue"`
}

type ResponseSalesAOV struct {
	StoreID uint64  `json:"store_id"`
	AOV     float64 `json:"aov"`
}

type ResponseProductSales struct {
	ProductID uint64  `json:"product_id"`
	Revenue   float64 `json:"revenue"`
	Quantity  float64 `json:"quantity"`
}

type ResponseSalesByProduct struct {
	StoreID  uint64                 `json:"store_id"`
	Products []ResponseProductSales `json:"products"`
}

type ResponseCategorySales struct {
	Category string  `json:"category"`
	Revenue  float64 `json:"revenue"`
	Quantity float64 `json:"quantity"`
}

type ResponseSalesByCategory struct {
	StoreID    uint64                  `json:"store_id"`
	Categories []ResponseCategorySales `json:"categories"`
}

type ResponseCustomerSales struct {
	CustomerID uint64  `json:"customer"`
	Revenue    float64 `json:"revenue"`
	Quantity   float64 `json:"quantity"`
}

type ResponseSalesCLV struct {
	StoreID uint64                  `json:"store_id"`
	CLV     []ResponseCustomerSales `json:"clv"`
}

func NewResponseSalesRevenue(c echo.Context, statusCode int, modelSale models.StoreSales) error {
	return Response(c, statusCode, ResponseSalesRevenue{
		StoreID: modelSale.StoreID,
		Revenue: modelSale.Price,
	})
}

func NewResponseSalesAOV(c echo.Context, statusCode int, modelSale models.StoreSales) error {
	return Response(c, statusCode, ResponseSalesAOV{
		StoreID: modelSale.StoreID,
		AOV:     modelSale.Price,
	})
}

func NewResponseSalesByProduct(c echo.Context, statusCode int, modelSales []models.ProductSales, storeID uint64) error {
	responseSales := make([]ResponseProductSales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseProductSales{
			ProductID: modelOrder.ProductID,
			Revenue:   modelOrder.Total,
			Quantity:  modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesByProduct{
		StoreID:  storeID,
		Products: responseSales,
	})
}

func NewResponseSalesByCategory(c echo.Context, statusCode int, modelSales []models.CategorySales, storeID uint64) error {
	responseSales := make([]ResponseCategorySales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseCategorySales{
			Category: modelOrder.Category,
			Revenue:  modelOrder.Total,
			Quantity: modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesByCategory{
		StoreID:    storeID,
		Categories: responseSales,
	})
}

func NewResponseSalesCLV(c echo.Context, statusCode int, modelSales []models.CustomerSales, storeID uint64) error {
	responseSales := make([]ResponseCustomerSales, 0)
	for _, modelOrder := range modelSales {
		responseSales = append(responseSales, ResponseCustomerSales{
			CustomerID: modelOrder.CustomerID,
			Revenue:    modelOrder.Total,
			Quantity:   modelOrder.Quantity,
		})
	}
	return Response(c, statusCode, ResponseSalesCLV{
		StoreID: storeID,
		CLV:     responseSales,
	})
}
