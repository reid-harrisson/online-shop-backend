package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseStore struct {
	ID                   uint64 `json:"id"`
	CompanyID            uint64 `json:"company_id"`
	OwnerID              uint64 `json:"owner_id"`
	ContactPhone         string `json:"contact_phone"`
	ContactEmail         string `json:"contact_email"`
	ShowStockLevelStatus string `json:"show_stock_level_status"`
	ShowOutOfStockStatus string `json:"show_out_of_stock_status"`
	BackOrderStatus      string `json:"back_order_status"`
	DeliveryPolicy       string `json:"delivery_policy"`
	ReturnsPolicy        string `json:"returns_policy"`
	Terms                string `json:"terms"`
}

type ResponseOutOfStockStatus struct {
	ShowOutOfStockStatus string `json:"show_out_of_stock_status"`
}

type ResponseStockLevelStatus struct {
	ShowStockLevelStatus string `json:"show_stock_level_status"`
}

type ResponseBackOrderStatus struct {
	BackOrderStatus string `json:"back_order_status"`
}

func NewResponseStore(c echo.Context, statusCode int, modelStore models.Stores) error {
	responseStore := ResponseStore{
		ID:                   uint64(modelStore.ID),
		CompanyID:            modelStore.CompanyID,
		OwnerID:              modelStore.OwnerID,
		ContactPhone:         modelStore.ContactPhone,
		ContactEmail:         modelStore.ContactEmail,
		ShowStockLevelStatus: utils.SimpleStatusToString(modelStore.ShowStockLevelStatus),
		ShowOutOfStockStatus: utils.SimpleStatusToString(modelStore.ShowOutOfStockStatus),
		BackOrderStatus:      utils.SimpleStatusToString(modelStore.BackOrderStatus),
		DeliveryPolicy:       modelStore.DeliveryPolicy,
		ReturnsPolicy:        modelStore.ReturnsPolicy,
		Terms:                modelStore.Terms,
	}
	return Response(c, statusCode, responseStore)
}

func NewResponseStores(c echo.Context, statusCode int, modelStores []models.Stores) error {
	responseStores := make([]ResponseStore, 0)
	for _, modelStore := range modelStores {
		responseStores = append(responseStores, ResponseStore{
			ID:                   uint64(modelStore.ID),
			CompanyID:            modelStore.CompanyID,
			OwnerID:              modelStore.OwnerID,
			ContactPhone:         modelStore.ContactPhone,
			ContactEmail:         modelStore.ContactEmail,
			ShowStockLevelStatus: utils.SimpleStatusToString(modelStore.ShowStockLevelStatus),
			ShowOutOfStockStatus: utils.SimpleStatusToString(modelStore.ShowOutOfStockStatus),
			BackOrderStatus:      utils.SimpleStatusToString(modelStore.BackOrderStatus),
			DeliveryPolicy:       modelStore.DeliveryPolicy,
			ReturnsPolicy:        modelStore.ReturnsPolicy,
			Terms:                modelStore.Terms,
		})
	}
	return Response(c, statusCode, responseStores)
}

func NewResponseOutOfStockStatus(c echo.Context, statusCode int, outOfStockStatus utils.SimpleStatuses) error {
	return Response(c, statusCode, ResponseOutOfStockStatus{
		ShowOutOfStockStatus: utils.SimpleStatusToString(outOfStockStatus),
	})
}

func NewResponseStockLevelStatus(c echo.Context, statusCode int, stockLevelStatus utils.SimpleStatuses) error {
	return Response(c, statusCode, ResponseStockLevelStatus{
		ShowStockLevelStatus: utils.SimpleStatusToString(stockLevelStatus),
	})
}

func NewResponseBackOrderStatus(c echo.Context, statusCode int, backOrderStatus utils.SimpleStatuses) error {
	return Response(c, statusCode, ResponseBackOrderStatus{
		BackOrderStatus: utils.SimpleStatusToString(backOrderStatus),
	})
}
