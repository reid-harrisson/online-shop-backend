package responses

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

	"github.com/labstack/echo/v4"
)

type ResponseVisitor struct {
	ID        uint64 `gorm:"column:id"`
	StoreID   uint64 `gorm:"column:store_id"`
	IpAddress string `gorm:"column:ip_address"`
	Page      string `gorm:"column:page"`
	Bounce    uint64 `gorm:"column:bounce"`
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
