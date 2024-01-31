package responses

import (
	"OnlineStoreBackend/models"
	"sort"

	"github.com/labstack/echo/v4"
)

type ResponseProductOrder struct {
	ProductID  uint64  `json:"store_product_id"`
	UnitPrice  float64 `json:"unit_price"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type ResponseStoreOrder struct {
	Product    []ResponseProductOrder `json:"product"`
	SubTotal   float64                `json:"sub_total"`
	TaxAmount  float64                `json:"tax_amount"`
	StoreTotal float64                `json:"store_total"`
}

type ResponseCustomerOrder struct {
	ID     uint64               `json:"id"`
	Store  []ResponseStoreOrder `json:"store"`
	Total  float64              `json:"total"`
	Status string               `json:"status"`
}

func NewResponseProductOrders(c echo.Context, statusCode int, modelOrders []models.ProductOrders) error {
	responseOrder := ResponseCustomerOrder{}
	sort.Slice(modelOrders, func(i, j int) bool {
		return modelOrders[i].StoreID > modelOrders[j].StoreID
	})
	storeID := uint64(0)
	responseOrder.Store = make([]ResponseStoreOrder, 0)
	responseStoreOrder := ResponseStoreOrder{}
	responseStoreOrder.Product = make([]ResponseProductOrder, 0)
	for _, modelOrder := range modelOrders {
		if storeID == 0 {
			storeID = modelOrder.StoreID
		} else if modelOrder.StoreID != storeID {
			responseStoreOrder.StoreTotal = responseStoreOrder.SubTotal + responseStoreOrder.TaxAmount
			responseOrder.Store = append(responseOrder.Store, responseStoreOrder)
			responseOrder.Total += responseStoreOrder.StoreTotal
			responseStoreOrder = ResponseStoreOrder{}
			responseStoreOrder.Product = make([]ResponseProductOrder, 0)
		}
		responseStoreOrder.SubTotal += modelOrder.SubTotal
		responseStoreOrder.TaxAmount += modelOrder.TaxAmount
		responseStoreOrder.Product = append(responseStoreOrder.Product, ResponseProductOrder{
			ProductID:  modelOrder.ProductID,
			UnitPrice:  modelOrder.UnitPriceSale,
			Quantity:   modelOrder.Quantity,
			TotalPrice: modelOrder.SubTotal,
		})
	}
	responseStoreOrder.StoreTotal = responseStoreOrder.SubTotal + responseStoreOrder.TaxAmount
	responseOrder.Store = append(responseOrder.Store, responseStoreOrder)
	responseOrder.ID = uint64(modelOrders[0].ID)
	responseOrder.Total += responseStoreOrder.StoreTotal
	responseOrder.Status = modelOrders[0].Status

	return Response(c, statusCode, responseOrder)
}
