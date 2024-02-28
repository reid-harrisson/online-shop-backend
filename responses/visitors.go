package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseVisitor struct {
	ID          uint64  `json:"id"`
	StoreID     uint64  `json:"store_id"`
	ProductID   uint64  `json:"product_id"`
	IpAddress   string  `json:"ip_address"`
	Page        string  `json:"page"`
	Bounce      uint64  `json:"bounce"`
	LoadingTime float64 `json:"lodaing_time"`
	ErrorCode   int     `json:"error_code"`
}

func NewResponseVisitor(c echo.Context, statusCode int, modelVisitor models.Visitors) error {
	return Response(c, statusCode, ResponseVisitor{
		ID:        uint64(modelVisitor.ID),
		StoreID:   modelVisitor.StoreID,
		IpAddress: modelVisitor.IpAddress,
		Page:      utils.PageTypeToString(modelVisitor.Page),
		Bounce:    modelVisitor.Bounce,
	})
}
