package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseStore struct {
	ID               uint64 `json:"id"`
	CompanyID        uint64 `json:"company_id"`
	OwnerID          uint64 `json:"owner_id"`
	ContactPhone     string `json:"contact_phone"`
	ContactEmail     string `json:"contact_email"`
	StockLevelStatus string `json:"show_stock_level_status"`
	OutOfStockStatus string `json:"show_out_of_stock_status"`
	BackOrderStatus  string `json:"back_order_status"`
	DeliveryPolicy   string `json:"delivery_policy"`
	ReturnsPolicy    string `json:"returns_policy"`
	Terms            string `json:"terms"`
}

type ResponseOutOfStockStatus struct {
	OutOfStockStatus string `json:"out_of_stock_status"`
}

type ResponseStockLevelStatus struct {
	StockLevelStatus string `json:"stock_level_status"`
}

type ResponseBackOrderStatus struct {
	BackOrderStatus string `json:"back_order_status"`
}

func NewResponseStore(c echo.Context, statusCode int, modelStore models.Stores) error {
	responseStore := ResponseStore{
		ID:               uint64(modelStore.ID),
		CompanyID:        modelStore.CompanyID,
		OwnerID:          modelStore.OwnerID,
		ContactPhone:     modelStore.ContactPhone,
		ContactEmail:     modelStore.ContactEmail,
		StockLevelStatus: utils.StockLevelStatusToString(modelStore.StockLevelStatus),
		OutOfStockStatus: utils.OutOfStockStatusToString(modelStore.OutOfStockStatus),
		BackOrderStatus:  utils.BackOrderStatusToString(modelStore.BackOrderStatus),
		DeliveryPolicy:   modelStore.DeliveryPolicy,
		ReturnsPolicy:    modelStore.ReturnsPolicy,
		Terms:            modelStore.Terms,
	}
	return Response(c, statusCode, responseStore)
}

func NewResponseOutOfStockStatus(c echo.Context, statusCode int, outOfStockStatus utils.OutOfStockStatus) error {
	return Response(c, statusCode, ResponseOutOfStockStatus{
		OutOfStockStatus: utils.OutOfStockStatusToString(outOfStockStatus),
	})
}

func NewResponseStockLevelStatus(c echo.Context, statusCode int, stockLevelStatus utils.StockLevelStatus) error {
	return Response(c, statusCode, ResponseStockLevelStatus{
		StockLevelStatus: utils.StockLevelStatusToString(stockLevelStatus),
	})
}

func NewResponseBackOrderStatus(c echo.Context, statusCode int, backOrderStatus utils.BackOrderStatus) error {
	return Response(c, statusCode, ResponseBackOrderStatus{
		BackOrderStatus: utils.BackOrderStatusToString(backOrderStatus),
	})
}
