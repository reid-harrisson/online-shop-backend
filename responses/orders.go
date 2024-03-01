package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseOrderItem struct {
	ID               uint64              `json:"id"`
	OrderID          uint64              `json:"order_id"`
	StoreID          uint64              `json:"store_id"`
	VariationID      uint64              `json:"variation_id"`
	Price            float64             `json:"price"`
	Quantity         float64             `json:"quantity"`
	SubTotalPrice    float64             `json:"sub_total_price"`
	TaxRate          float64             `json:"tax_rate"`
	TaxAmount        float64             `json:"tax_amount"`
	ShippingMethodID uint64              `json:"shipping_method_id"`
	ShippingPrice    float64             `json:"shipping_price"`
	TotalPrice       float64             `json:"total_price"`
	Status           utils.OrderStatuses `json:"status"`
}

type ResponseCustomerOrderItem struct {
	StoreID          uint64  `json:"store_id"`
	ProductStatus    string  `json:"product_status"`
	ProductID        uint64  `json:"product_id"`
	Price            float64 `json:"price"`
	Quantity         float64 `json:"quantity"`
	SubTotalPrice    float64 `json:"sub_total_price"`
	TaxRate          float64 `json:"tax_rate"`
	TaxAmount        float64 `json:"tax_amount"`
	ShippingMethodID uint64  `json:"shipping_method_id"`
	ShippingPrice    float64 `json:"shipping_price"`
	TotalPrice       float64 `json:"total_price"`
}

type ResponseCustomerOrderWithDetail struct {
	OrderStatus     string                      `json:"order_status"`
	ShippingAddress ResponseCustomerAddress     `json:"shipping_address"`
	BillingAddress  ResponseCustomerAddress     `json:"billing_address"`
	Products        []ResponseCustomerOrderItem `json:"items"`
}

type ResponseCustomerOrder struct {
	OrderID           uint64  `json:"order_id"`
	OrderStatus       string  `json:"order_status"`
	TotalPrice        float64 `json:"total_price"`
	BillingAddressID  uint64  `json:"billing_address"`
	ShippingAddressID uint64  `json:"shipping_address"`
}

type ResponseStoreOrder struct {
	OrderID           uint64  `json:"order_id"`
	CustomerID        uint64  `json:"customer_id"`
	VariationID       uint64  `json:"variation_id"`
	Price             float64 `json:"price"`
	Quantity          float64 `json:"quantity"`
	SubTotalPrice     float64 `json:"sub_total_price"`
	BillingAddressID  uint64  `json:"billing_address"`
	ShippingAddressID uint64  `json:"shipping_address"`
	TaxRate           float64 `json:"tax_rate"`
	TaxAmount         float64 `json:"tax_amount"`
	ShippingMethodID  uint64  `json:"shipping_method_id"`
	ShippingPrice     float64 `json:"shipping_price"`
	TotalPrice        float64 `json:"total_price"`
	ProductStatus     string  `json:"product_status"`
}

func NewResponseCustomerOrders(c echo.Context, statusCode int, modelOrders []models.CustomerOrders) error {
	responseOrders := make([]ResponseCustomerOrder, 0)
	for _, modelOrder := range modelOrders {
		responseOrders = append(responseOrders, ResponseCustomerOrder{
			OrderID:           modelOrder.OrderID,
			OrderStatus:       utils.OrderStatusToString(modelOrder.OrderStatus),
			TotalPrice:        modelOrder.TotalPrice,
			BillingAddressID:  modelOrder.BillingAddressID,
			ShippingAddressID: modelOrder.ShippingAddressID,
		})
	}
	return Response(c, statusCode, responseOrders)
}

func NewResponseCustomerOrdersWithDetail(c echo.Context, statusCode int, modelOrder models.CustomerOrdersWithAddress) error {
	responseItems := make([]ResponseCustomerOrderItem, 0)
	orderStatus := utils.StatusOrderPending
	for _, modelItem := range modelOrder.Items {
		responseItems = append(responseItems, ResponseCustomerOrderItem{
			StoreID:          modelItem.StoreID,
			ProductStatus:    utils.OrderStatusToString(modelItem.ProductStatus),
			ProductID:        modelItem.VariationID,
			Price:            modelItem.Price,
			Quantity:         modelItem.Quantity,
			SubTotalPrice:    modelItem.SubTotalPrice,
			TaxRate:          modelItem.TaxRate,
			TaxAmount:        modelItem.TaxAmount,
			ShippingMethodID: modelItem.ShippingMethodID,
			ShippingPrice:    modelItem.ShippingPrice,
			TotalPrice:       modelItem.TotalPrice,
		})
		orderStatus = modelItem.OrderStatus
	}
	return Response(c, statusCode, ResponseCustomerOrderWithDetail{
		OrderStatus: utils.OrderStatusToString(orderStatus),
		BillingAddress: ResponseCustomerAddress{
			ID:           uint64(modelOrder.BillingAddress.ID),
			AddressLine1: modelOrder.BillingAddress.AddressLine1,
			AddressLine2: modelOrder.BillingAddress.AddressLine2,
			SubUrb:       modelOrder.BillingAddress.SubUrb,
			CountryID:    modelOrder.BillingAddress.CountryID,
			RegionID:     modelOrder.BillingAddress.RegionID,
			CityID:       modelOrder.BillingAddress.CityID,
			PostalCode:   modelOrder.BillingAddress.PostalCode,
		},
		ShippingAddress: ResponseCustomerAddress{
			ID:           uint64(modelOrder.ShippingAddress.ID),
			AddressLine1: modelOrder.ShippingAddress.AddressLine1,
			AddressLine2: modelOrder.ShippingAddress.AddressLine2,
			SubUrb:       modelOrder.ShippingAddress.SubUrb,
			CountryID:    modelOrder.ShippingAddress.CountryID,
			RegionID:     modelOrder.ShippingAddress.RegionID,
			CityID:       modelOrder.ShippingAddress.CityID,
			PostalCode:   modelOrder.ShippingAddress.PostalCode,
		},
		Products: responseItems,
	})
}

func NewResponseStoreOrders(c echo.Context, statusCode int, modelOrders []models.StoreOrders) error {
	responseOrders := make([]ResponseStoreOrder, 0)
	for _, modelOrder := range modelOrders {
		responseOrders = append(responseOrders, ResponseStoreOrder{
			OrderID:           modelOrder.OrderID,
			CustomerID:        modelOrder.CustomerID,
			VariationID:       modelOrder.VariationID,
			Price:             modelOrder.Price,
			Quantity:          modelOrder.Quantity,
			SubTotalPrice:     modelOrder.SubTotalPrice,
			BillingAddressID:  modelOrder.BillingAddressID,
			ShippingAddressID: modelOrder.ShippingAddressID,
			TaxRate:           modelOrder.TaxRate,
			TaxAmount:         modelOrder.TaxAmount,
			ShippingMethodID:  modelOrder.ShippingMethodID,
			ShippingPrice:     modelOrder.ShippingPrice,
			TotalPrice:        modelOrder.TotalPrice,
			ProductStatus:     utils.OrderStatusToString(modelOrder.ProductStatus),
		})
	}
	return Response(c, statusCode, responseOrders)
}

func NewResponseStoreOrder(c echo.Context, statusCode int, modelOrder models.StoreOrders) error {
	return Response(c, statusCode, ResponseStoreOrder{
		OrderID:           modelOrder.OrderID,
		CustomerID:        modelOrder.CustomerID,
		VariationID:       modelOrder.VariationID,
		Price:             modelOrder.Price,
		Quantity:          modelOrder.Quantity,
		SubTotalPrice:     modelOrder.SubTotalPrice,
		BillingAddressID:  modelOrder.BillingAddressID,
		ShippingAddressID: modelOrder.ShippingAddressID,
		TaxRate:           modelOrder.TaxRate,
		TaxAmount:         modelOrder.TaxAmount,
		ShippingMethodID:  modelOrder.ShippingMethodID,
		ShippingPrice:     modelOrder.ShippingPrice,
		TotalPrice:        modelOrder.TotalPrice,
		ProductStatus:     utils.OrderStatusToString(modelOrder.ProductStatus),
	})
}

func NewResponseOrderItems(c echo.Context, statusCode int, modelItems []models.OrderItems) error {
	responseItems := make([]ResponseOrderItem, 0)
	for _, modelItem := range modelItems {
		responseItems = append(responseItems, ResponseOrderItem{
			ID:               uint64(modelItem.ID),
			OrderID:          modelItem.OrderID,
			StoreID:          modelItem.StoreID,
			VariationID:      modelItem.VariationID,
			Price:            modelItem.Price,
			Quantity:         modelItem.Quantity,
			SubTotalPrice:    modelItem.SubTotalPrice,
			TaxRate:          modelItem.TaxRate,
			TaxAmount:        modelItem.TaxAmount,
			ShippingMethodID: modelItem.ShippingMethodID,
			ShippingPrice:    modelItem.ShippingPrice,
			TotalPrice:       modelItem.TotalPrice,
			Status:           modelItem.Status,
		})
	}
	return Response(c, statusCode, responseItems)
}
