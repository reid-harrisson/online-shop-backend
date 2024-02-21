package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseCustomerOrderItem struct {
	StoreID        uint64  `json:"store_id"`
	ProductStatus  string  `json:"product_status"`
	ProductID      uint64  `json:"product_id"`
	UnitPriceSale  float64 `json:"unit_price_sale"`
	Quantity       float64 `json:"quantity"`
	SubTotalPrice  float64 `json:"sub_total_price"`
	TaxRate        float64 `json:"tax_rate"`
	TaxAmount      float64 `json:"tax_amount"`
	ShippingMethod string  `json:"shipping_method"`
	ShippingPrice  float64 `json:"shipping_price"`
	TotalPrice     float64 `json:"total_price"`
}

type ResponseCustomerOrderWithDetail struct {
	OrderStatus     string                      `json:"order_status"`
	ShippingAddress ResponseCustomerAddress     `json:"shipping_address"`
	BillingAddress  ResponseCustomerAddress     `json:"billing_address"`
	Items           []ResponseCustomerOrderItem `json:"items"`
}

type ResponseCustomerOrder struct {
	OrderID           uint64  `json:"order_id"`
	OrderStatus       string  `json:"order_status"`
	TotalPrice        float64 `json:"total_price"`
	BillingAddressID  uint64  `json:"billing_address"`
	ShippingAddressID uint64  `json:"shipping_address"`
}

type ResponseStoreOrder struct {
	OrderID           uint64  `gorm:"column:order_id"`
	CustomerID        uint64  `gorm:"column:customer_id"`
	ProductID         uint64  `gorm:"column:product_id"`
	UnitPriceSale     float64 `gorm:"column:unit_price_sale"`
	Quantity          float64 `gorm:"column:quantity"`
	SubTotalPrice     float64 `gorm:"column:sub_total_price"`
	BillingAddressID  uint64  `gorm:"column:billing_address"`
	ShippingAddressID uint64  `gorm:"column:shipping_address"`
	TaxRate           float64 `gorm:"column:tax_rate"`
	TaxAmount         float64 `gorm:"column:tax_amount"`
	ShippingMethod    string  `gorm:"column:shipping_method"`
	ShippingPrice     float64 `gorm:"column:shipping_price"`
	TotalPrice        float64 `gorm:"column:total_price"`
	ProductStatus     string  `gorm:"column:product_status"`
}

func NewResponseCustomerOrders(c echo.Context, statusCode int, modelOrders []models.CustomerOrders) error {
	responseOrders := make([]ResponseCustomerOrder, 0)
	for _, modelOrder := range modelOrders {
		responseOrders = append(responseOrders, ResponseCustomerOrder{
			OrderID:           modelOrder.OrderID,
			OrderStatus:       models.OrderStatusToString(modelOrder.OrderStatus),
			TotalPrice:        modelOrder.TotalPrice,
			BillingAddressID:  modelOrder.BillingAddressID,
			ShippingAddressID: modelOrder.ShippingAddressID,
		})
	}
	return Response(c, statusCode, responseOrders)
}

func NewResponseCustomerOrdersWithDetail(c echo.Context, statusCode int, modelOrders []models.CustomerOrdersWithDetail) error {
	responseItems := make([]ResponseCustomerOrderItem, 0)
	orderStatus := models.StatusOrderPending
	billingAddress := new(models.CustomerAddresses)
	shippingAddress := new(models.CustomerAddresses)
	for _, modelOrder := range modelOrders {
		if orderStatus == 0 {
			orderStatus = modelOrder.OrderStatus
			billingAddress = modelOrder.BillingAddress
			shippingAddress = modelOrder.ShippingAddress
		}
		responseItems = append(responseItems, ResponseCustomerOrderItem{
			StoreID:        modelOrder.StoreID,
			ProductStatus:  models.OrderStatusToString(modelOrder.ProductStatus),
			ProductID:      modelOrder.ProductID,
			UnitPriceSale:  modelOrder.UnitPriceSale,
			Quantity:       modelOrder.Quantity,
			SubTotalPrice:  modelOrder.SubTotalPrice,
			TaxRate:        modelOrder.TaxRate,
			TaxAmount:      modelOrder.TaxAmount,
			ShippingMethod: modelOrder.ShippingMethod,
			ShippingPrice:  modelOrder.ShippingPrice,
			TotalPrice:     modelOrder.TotalPrice,
		})
	}
	return Response(c, statusCode, ResponseCustomerOrderWithDetail{
		OrderStatus: models.OrderStatusToString(models.OrderStatuses(orderStatus)),
		BillingAddress: ResponseCustomerAddress{
			ID:           uint64(billingAddress.ID),
			AddressLine1: billingAddress.AddressLine1,
			AddressLine2: billingAddress.AddressLine2,
			SubUrb:       billingAddress.SubUrb,
			CountryID:    billingAddress.CountryID,
			RegionID:     billingAddress.RegionID,
			CityID:       billingAddress.CityID,
			PostalCode:   billingAddress.PostalCode,
		},
		ShippingAddress: ResponseCustomerAddress{
			ID:           uint64(shippingAddress.ID),
			AddressLine1: shippingAddress.AddressLine1,
			AddressLine2: shippingAddress.AddressLine2,
			SubUrb:       shippingAddress.SubUrb,
			CountryID:    shippingAddress.CountryID,
			RegionID:     shippingAddress.RegionID,
			CityID:       shippingAddress.CityID,
			PostalCode:   shippingAddress.PostalCode,
		},
		Items: responseItems,
	})
}

func NewResponseStoreOrders(c echo.Context, statusCode int, modelOrders []models.StoreOrders) error {
	responseOrders := make([]ResponseStoreOrder, 0)
	for _, modelOrder := range modelOrders {
		responseOrders = append(responseOrders, ResponseStoreOrder{
			OrderID:           modelOrder.OrderID,
			CustomerID:        modelOrder.CustomerID,
			ProductID:         modelOrder.ProductID,
			UnitPriceSale:     modelOrder.UnitPriceSale,
			Quantity:          modelOrder.Quantity,
			SubTotalPrice:     modelOrder.SubTotalPrice,
			BillingAddressID:  modelOrder.BillingAddressID,
			ShippingAddressID: modelOrder.ShippingAddressID,
			TaxRate:           modelOrder.TaxRate,
			TaxAmount:         modelOrder.TaxAmount,
			ShippingMethod:    modelOrder.ShippingMethod,
			ShippingPrice:     modelOrder.ShippingPrice,
			TotalPrice:        modelOrder.TotalPrice,
			ProductStatus:     models.OrderStatusToString(modelOrder.ProductStatus),
		})
	}
	return Response(c, statusCode, responseOrders)
}

func NewResponseStoreOrder(c echo.Context, statusCode int, modelOrder models.StoreOrders) error {
	return Response(c, statusCode, ResponseStoreOrder{
		OrderID:           modelOrder.OrderID,
		CustomerID:        modelOrder.CustomerID,
		ProductID:         modelOrder.ProductID,
		UnitPriceSale:     modelOrder.UnitPriceSale,
		Quantity:          modelOrder.Quantity,
		SubTotalPrice:     modelOrder.SubTotalPrice,
		BillingAddressID:  modelOrder.BillingAddressID,
		ShippingAddressID: modelOrder.ShippingAddressID,
		TaxRate:           modelOrder.TaxRate,
		TaxAmount:         modelOrder.TaxAmount,
		ShippingMethod:    modelOrder.ShippingMethod,
		ShippingPrice:     modelOrder.ShippingPrice,
		TotalPrice:        modelOrder.TotalPrice,
		ProductStatus:     models.OrderStatusToString(modelOrder.ProductStatus),
	})
}
