package responses

import (
	"OnlineStoreBackend/models"

	"github.com/labstack/echo/v4"
)

type ResponseStore struct {
	ID                   uint64 `json:"id"`
	CompanyID            uint64 `json:"company_id"`
	OwnerID              uint64 `json:"owner_id"`
	ContactPhone         string `json:"contact_phone"`
	ContactEmail         string `json:"contact_email"`
	ShowStockLevelStatus int8   `json:"show_stock_level_status"`
	ShowOutOfStockStatus int8   `json:"show_out_of_stock_status"`
	IsBackOrder          int8   `json:"is_back_order"`
	DeliveryPolicy       string `json:"delivery_policy"`
	ReturnsPolicy        string `json:"returns_policy"`
	Terms                string `json:"terms"`
	Active               int8   `json:"active"`
}

type ResponseShowOutOfStockStatus struct {
	ShowOutOfStockStatus string `json:"show_out_of_stock_status"`
}

type ResponseShowStockLevelStatus struct {
	ShowStockLevelStatus string `json:"show_stock_level_status"`
}

type ResponseBackOrder struct {
	IsBackOrder string `json:"back_order"`
}

func NewResponseStore(c echo.Context, statusCode int, modelStore models.Stores) error {
	responseStore := ResponseStore{
		ID:                   uint64(modelStore.ID),
		CompanyID:            modelStore.CompanyID,
		OwnerID:              modelStore.OwnerID,
		ContactPhone:         modelStore.ContactPhone,
		ContactEmail:         modelStore.ContactEmail,
		ShowStockLevelStatus: modelStore.ShowStockLevelStatus,
		ShowOutOfStockStatus: modelStore.ShowOutOfStockStatus,
		IsBackOrder:          modelStore.IsBackOrder,
		DeliveryPolicy:       modelStore.DeliveryPolicy,
		ReturnsPolicy:        modelStore.ReturnsPolicy,
		Terms:                modelStore.Terms,
		Active:               modelStore.Active,
	}
	return Response(c, statusCode, responseStore)
}

func NewResponseShowOutOfStockStatus(c echo.Context, statusCode int, isShowOutOfStockStatus int8) error {
	if isShowOutOfStockStatus == 0 {
		return Response(c, statusCode, ResponseShowOutOfStockStatus{
			ShowOutOfStockStatus: "Disabled",
		})
	}
	return Response(c, statusCode, ResponseShowOutOfStockStatus{
		ShowOutOfStockStatus: "Enabled",
	})
}

func NewResponseShowStockLevelStatus(c echo.Context, statusCode int, isShowStockLevelStatus int8) error {
	if isShowStockLevelStatus == 0 {
		return Response(c, statusCode, ResponseShowStockLevelStatus{
			ShowStockLevelStatus: "Disabled",
		})
	}
	return Response(c, statusCode, ResponseShowStockLevelStatus{
		ShowStockLevelStatus: "Enabled",
	})
}

func NewResponseBackOrder(c echo.Context, statusCode int, backOrder int8) error {
	if backOrder == 0 {
		return Response(c, statusCode, ResponseBackOrder{
			IsBackOrder: "Disabled",
		})
	}
	return Response(c, statusCode, ResponseBackOrder{
		IsBackOrder: "Enabled",
	})
}
