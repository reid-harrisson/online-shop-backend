package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseStore struct {
	ID                   uint64               `json:"id" example:"2"`
	CompanyID            uint64               `json:"company_id" example:"2"`
	OwnerID              uint64               `json:"owner_id" example:"1607"`
	Name                 string               `json:"name" example:"The Che Gourmet Shop"`
	ContactPhone         string               `json:"contact_phone" example:"7184756027"`
	ContactEmail         string               `json:"contact_email" example:"example@sample.com"`
	ShowStockLevelStatus utils.SimpleStatuses `json:"show_stock_level_status" example:"0"`
	ShowOutOfStockStatus utils.SimpleStatuses `json:"show_out_of_stock_status" example:"0"`
	DeliveryPolicy       string               `json:"delivery_policy" example:"example delivery policy"`
	ReturnsPolicy        string               `json:"returns_policy" example:"example return policy"`
	Terms                string               `json:"terms" example:"example terms"`
}

type ResponseOutOfStockStatus struct {
	ShowOutOfStockStatus utils.SimpleStatuses `json:"show_out_of_stock_status" example:"0"`
}

type ResponseStockLevelStatus struct {
	ShowStockLevelStatus utils.SimpleStatuses `json:"show_stock_level_status" example:"0"`
}

type ResponseBackOrderStatus struct {
	BackOrderStatus string `json:"back_order_status"`
}

type ResponseMinimumStockLevel struct {
	MinimumStockLevel float64 `json:"minimum_stock_level"`
}

type ResponseStockLevel struct {
	StockLevel float64 `json:"stock_level"`
}

func NewResponseStore(c echo.Context, statusCode int, modelStore models.Stores) error {
	responseStore := ResponseStore{
		ID:                   uint64(modelStore.ID),
		CompanyID:            modelStore.CompanyID,
		OwnerID:              modelStore.OwnerID,
		Name:                 modelStore.Name,
		ContactPhone:         modelStore.ContactPhone,
		ContactEmail:         modelStore.ContactEmail,
		ShowStockLevelStatus: modelStore.ShowStockLevelStatus,
		ShowOutOfStockStatus: modelStore.ShowOutOfStockStatus,
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
			Name:                 modelStore.Name,
			ContactPhone:         modelStore.ContactPhone,
			ContactEmail:         modelStore.ContactEmail,
			ShowStockLevelStatus: modelStore.ShowStockLevelStatus,
			ShowOutOfStockStatus: modelStore.ShowOutOfStockStatus,
			DeliveryPolicy:       modelStore.DeliveryPolicy,
			ReturnsPolicy:        modelStore.ReturnsPolicy,
			Terms:                modelStore.Terms,
		})
	}
	return Response(c, statusCode, responseStores)
}

func NewResponseOutOfStockStatus(c echo.Context, statusCode int, outOfStockStatus utils.SimpleStatuses) error {
	return Response(c, statusCode, ResponseOutOfStockStatus{
		ShowOutOfStockStatus: outOfStockStatus,
	})
}

func NewResponseStockLevelStatus(c echo.Context, statusCode int, stockLevelStatus utils.SimpleStatuses) error {
	return Response(c, statusCode, ResponseStockLevelStatus{
		ShowStockLevelStatus: stockLevelStatus,
	})
}

func NewResponseMinimumStockLevel(c echo.Context, statusCode int, minimumStockLevel float64) error {
	return Response(c, statusCode, ResponseMinimumStockLevel{
		MinimumStockLevel: minimumStockLevel,
	})
}

func NewResponseStockLevel(c echo.Context, statusCode int, stockLevel float64) error {
	return Response(c, statusCode, ResponseStockLevel{
		StockLevel: stockLevel,
	})
}
